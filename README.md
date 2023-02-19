# My Common Docker Images

This is a small collection of docker image recipes that I reuse across several personal proejcts.

## Getting Started

1. Build local tools.

```
$ make
```

2. Use `bin/build` to build any of the images. For instance, to build `kellegous/build`, run the following:

```
$ bin/build build
```

This will build the image declared in the "build" sub-directory and export the results into tarballs for both `linux/amd64` and `linux/arm64`. To push those images to [hub.docker.com](https://hub.docker.com/) instead of exporting them, run the following:

```
$ bin/build --push build
```

