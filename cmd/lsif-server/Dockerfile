
# TODO: Make this image use our sourcegraph/alpine:3.9 base image (3.10?).
FROM alpine:3.10@sha256:e4355b66995c96b4b468159fc5c7e3540fcef961189ca13fee877798649f531a

ARG COMMIT_SHA="unknown"
ARG DATE="unknown"
ARG VERSION="unknown"

LABEL org.opencontainers.image.revision=${COMMIT_SHA}
LABEL org.opencontainers.image.created=${DATE}
LABEL org.opencontainers.image.version=${VERSION}
LABEL com.sourcegraph.github.url=https://github.com/sourcegraph/sourcegraph/commit/${COMMIT_SHA}

# set default environment variables
ENV PGDATABASE=sg PGHOST=pgsql PGPORT=5432 PGSSLMODE=disable PGUSER=sg

# hadolint ignore=DL3018
RUN apk update && apk add --no-cache \
    tini nodejs-current=12.4.0-r0

# hadolint ignore=DL3022
COPY --from=sourcegraph/lsif-server-builder:ci /lsif /lsif

COPY . /

EXPOSE 3186 3187
ENV GO111MODULES=on LANG=en_US.utf8
ENTRYPOINT ["/sbin/tini", "--", "/usr/local/bin/lsif-server"]
