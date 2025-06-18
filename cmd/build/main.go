package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"

	"github.com/spf13/cobra"
)

func buildImage(
	ctx context.Context,
	dockerFile string,
	name string,
	tag string,
	push bool,
) error {
	args := []string{
		fmt.Sprintf("--dockerfile=%s", dockerFile),
		fmt.Sprintf("--root=%s", filepath.Dir(dockerFile)),
	}

	if tag != "" {
		args = append(args, fmt.Sprintf("--tag=%s", tag))
	}

	if push {
		args = append(
			args,
			"--target=linux/amd64",
			"--target=linux/arm64")
	} else {
		args = append(args,
			fmt.Sprintf("--target=linux/amd64:%s-amd64.tar", name),
			fmt.Sprintf("--target=linux/arm64:%s-arm64.tar", name))
	}

	args = append(args, fmt.Sprintf("kellegous/%s", name))

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

			name := args[0]
			dfPath := filepath.Join(name, "Dockerfile")
			if _, err := os.Stat(dfPath); err != nil {
				cmd.PrintErrf("file not found: %s", dfPath)
				os.Exit(1)
			}

			if err := buildImage(ctx, dfPath, name, tag, push); err != nil {
				cmd.PrintErrf("build image: %s", err)
				os.Exit(1)
			}
		},
	}

	c.Flags().BoolVar(
		&push,
		"push",
		false,
		"whether to push the image to ghcr.io")

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
