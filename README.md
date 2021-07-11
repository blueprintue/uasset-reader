## Uasset Reader

[![GitHub release](https://img.shields.io/github/release/blueprintue/uasset-reader.svg)](https://github.com/blueprintue/uasset-reader/releases/latest)
[![Total downloads](https://img.shields.io/github/downloads/blueprintue/uasset-reader/total.svg)](https://github.com/blueprintue/uasset-reader/releases/latest)
[![Build Status](https://img.shields.io/github/workflow/status/blueprintue/uasset-reader/build?label=build&logo=github)](https://github.com/blueprintue/uasset-reader/actions?query=workflow%3Abuild)
[![Docker Stars](https://img.shields.io/docker/stars/blueprintue/uasset-reader.svg&logo=docker)](https://hub.docker.com/r/blueprintue/uasset-reader/)
[![Docker Pulls](https://img.shields.io/docker/pulls/blueprintue/uasset-reader.svg&logo=docker)](https://hub.docker.com/r/blueprintue/uasset-reader/)
[![Go Report Card](https://goreportcard.com/badge/github.com/blueprintue/uasset-reader)](https://goreportcard.com/report/github.com/blueprintue/uasset-reader)

## Usage

### From binary

`uasset-reader` binaries are available on [releases page](https://github.com/blueprintue/uasset-reader/releases/latest).

Choose the archive matching the destination platform:

```shell
wget -qO- https://github.com/blueprintue/uasset-reader/releases/download/v0.1.0/uasset-reader_0.1.0_linux_arm64.tar.gz | tar -zxvf - uasset-reader
```

### From Dockerfile

| Registry                                                                                                    | Image                             |
|-------------------------------------------------------------------------------------------------------------|-----------------------------------|
| [Docker Hub](https://hub.docker.com/r/blueprintue/uasset-reader/)                                           | `blueprintue/uasset-reader`         |
| [GitHub Container Registry](https://github.com/users/blueprintue/packages/container/package/uasset-reader)  | `ghcr.io/blueprintue/uasset-reader` |

Following platforms for this image are available:

#todo: to update

```
$ docker buildx imagetools inspect blueprintue/uasset-reader:latest
Name:      docker.io/blueprintue/uasset-reader:edge
MediaType: application/vnd.docker.distribution.manifest.list.v2+json
Digest:    sha256:8d4501ea22b8914315a26acbf3da17c1009d15e480a80a75f1d2166db48ac998

Manifests:
  Name:      docker.io/blueprintue/uasset-reader:edge@sha256:7194ed60cdcd51c57389c666d7e325464486e6fb7f593c3db57230ff0f05c40b
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/amd64

  Name:      docker.io/blueprintue/uasset-reader:edge@sha256:7c984e4b0a0f9e106d201e15633e17319e4a312d283a16fa7b2782b6ddc9bb57
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/arm/v7

  Name:      docker.io/blueprintue/uasset-reader:edge@sha256:329c01304b303d64016ceffdfa1b1c50731212d72db5e070f04923f0375f4df4
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/arm64
```

## Build

```shell
git clone https://github.com/blueprintue/uasset-reader.git uasset-reader
cd uasset-reader

# build docker image and output to docker with uasset-reader tag (default)
docker buildx bake

# build multi-platform image
docker buildx bake image-all

# create artifacts in ./dist
docker buildx bake artifact-all
```