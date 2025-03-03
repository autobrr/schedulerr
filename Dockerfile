# build app
FROM golang:1.23-alpine3.20 AS app-builder

ARG VERSION=dev
ARG REVISION=dev
ARG BUILDTIME

RUN apk add --no-cache git make build-base

ENV SERVICE=schedulerr

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

#ENV GOOS=linux
ENV CGO_ENABLED=0

RUN go build -ldflags "-s -w -X github.com/autobrr/schedulerr/internal/buildinfo.Version=${VERSION} -X github.com/autobrr/schedulerr/internal/buildinfo.Commit=${REVISION} -X github.com/autobrr/schedulerr/internal/buildinfo.Date=${BUILDTIME}" -o bin/schedulerr cmd/schedulerr/main.go

# build runner
FROM alpine:latest

LABEL org.opencontainers.image.source="https://github.com/autobrr/schedulerr"

ENV APP_DIR="/app" CONFIG_DIR="/config" PUID="1000" PGID="1000" UMASK="002" TZ="Etc/UTC" ARGS=""
ENV XDG_CONFIG_HOME="${CONFIG_DIR}/.config" XDG_CACHE_HOME="${CONFIG_DIR}/.cache" XDG_DATA_HOME="${CONFIG_DIR}/.local/share" LANG="C.UTF-8" LC_ALL="C.UTF-8"

VOLUME ["${CONFIG_DIR}"]

# install packages
RUN apk add --no-cache tzdata shadow bash curl wget jq grep sed coreutils findutils unzip p7zip ca-certificates

COPY --from=app-builder /src/bin/schedulerr /usr/bin/

# make folders
RUN mkdir "${APP_DIR}" && \
    # create user
    useradd -u 1000 -U -d "${CONFIG_DIR}" -s /bin/false schedulerr && \
    usermod -G users schedulerr

WORKDIR /config

EXPOSE 8585

ENTRYPOINT ["schedulerr", "--config", "/config/config.yaml"]
CMD ["serve"]