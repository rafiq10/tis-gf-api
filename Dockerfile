# syntax = docker/dockerfile:1.2

# FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base
FROM  golang:1.16-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0 
ENV PATH=$PATH:/opt/mssql-tools/bin
COPY go.* ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS 
ARG TARGETARCH 
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/main .

FROM base AS unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir /test && go test -v -coverprofile=/test/cover.out ./...

FROM golangci/golangci-lint:v1.38.0-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 10m0s ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /bin/cover.out /cover.out

FROM scratch AS bin-unix
COPY --from=build /bin/main /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /bin/main /main.exe

FROM bin-${TARGETOS} as bin