# syntax=docker/dockerfile:1.2

FROM --platform=$BUILDPLATFORM crazymax/goreleaser-xx:latest AS goreleaser-xx
FROM --platform=$BUILDPLATFORM golang:alpine AS base
COPY --from=goreleaser-xx / /
RUN apk add --no-cache git
WORKDIR /src

FROM base AS build
ARG TARGETPLATFORM
ARG GIT_REF
RUN --mount=type=bind,source=.,target=/src,rw \
  goreleaser-xx --debug \
    --name="uasset-reader" \
    --dist="/out" \
    --hooks="go mod tidy" \
    --hooks="go mod download" \
    --ldflags="-s -w -X 'main.version={{.Version}}'" \
    --files="LICENSE" \
    --files="README.md"

FROM scratch AS artifact
COPY --from=build /out/*.tar.gz /
COPY --from=build /out/*.zip /

FROM alpine:3.13 AS image
RUN apk --update --no-cache add ca-certificates libressl shadow \
  && addgroup -g 1000 uasset-reader \
  && adduser -u 1000 -G uasset-reader -s /sbin/nologin -D uasset-reader \
  && mkdir -p /var/log/uasset-reader \
  && chown uasset-reader. /var/log/uasset-reader
COPY --from=build /usr/local/bin/uasset-reader /usr/local/bin/uasset-reader
USER uasset-reader
ENTRYPOINT [ "uasset-reader" ]