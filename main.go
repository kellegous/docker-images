package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func getTagFromGit(
	ctx context.Context,
) (string, error) {
	var buf bytes.Buffer
	c := exec.CommandContext(ctx, "git", "rev-parse", "HEAD")
	c.Stderr = os.Stderr
	c.Stdout = &buf
	if err := c.Run(); err != nil {
		return "", err
	}
	tag := strings.TrimSpace(buf.String())
	if len(tag) > 8 {
		tag = tag[:8]
	}
	return tag, nil
}

func buildImage(
	ctx context.Context,
	dockerFile string,
	name string,
	tag string,
	push bool,
) error {
	args := []string{
		fmt.Sprintf("--docker-file=%s", dockerFile),
		fmt.Sprintf("--version=%s", tag),
		"--platform=linux/arm64",
		"--platform=linux/amd64",
	}

	if push {
		args = append(args, "--push")
	}

	args = append(args, name)

	c := exec.CommandContext(ctx, "bin/buildimg", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	return c.Run()
}

func buildCommand() *cobra.Command {
	var push bool
	var tag string

	c := &cobra.Command{
		Use:   "build [flags] image",
		Short: "build and/or push the images declared in this repo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, done := signal.NotifyContext(
				context.Background(),
				os.Interrupt)
			defer done()

			var err error
			if tag == "" {
				tag, err = getTagFromGit(ctx)
				if err != nil {
					cmd.PrintErrf("git rev-parse: %s", err)
					os.Exit(1)
				}
			}

			name := args[0]
			dfPath := filepath.Join(name, "Dockerfile")
			if _, err := os.Stat(dfPath); err != nil {
				cmd.PrintErrf("file not found: %s", dfPath)
				os.Exit(1)
			}

			image := fmt.Sprintf("kellegous/%s", name)
			if err := buildImage(ctx, dfPath, image, tag, push); err != nil {
				cmd.PrintErrf("build image: %s", err)
				os.Exit(1)
			}

			fmt.Printf("%s:%s\n", image, tag)
		},
	}

	c.Flags().BoolVar(
		&push,
		"push",
		false,
		"whether to push the image to hub.docker.com")

	c.Flags().StringVar(
		&tag,
		"tag",
		"",
		"the image tag (default based on git SHA)")

	return c
}

func main() {
	if err := buildCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
