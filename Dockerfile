# Stage 1: Build the Go project
FROM golang:1.23-alpine3.20 AS go-builder

# Use build arguments for the target architecture
ARG TARGETARCH
ARG GOARCH
ARG VERSION
ARG COMMIT

ENV MIMALLOC_VERSION=v2.2.2

# Install necessary packages
RUN set -eux; apk add --no-cache ca-certificates build-base git cmake

WORKDIR /code
COPY . /code/

# Install mimalloc
RUN git clone -b ${MIMALLOC_VERSION} --depth 1 https://github.com/microsoft/mimalloc; cd mimalloc; mkdir build; cd build; cmake ..; make -j$(nproc); make install
ENV MIMALLOC_RESERVE_HUGE_OS_PAGES=4

RUN VERSION=${VERSION} COMMIT=${COMMIT} LEDGER_ENABLED=false GOARCH=${GOARCH} LDFLAGS="-linkmode=external -extldflags \"-L/code/mimalloc/build -lmimalloc -Wl,-z,muldefs -static\"" make build

# use bullseye-slim as base image for rly binary at launch
FROM debian:bullseye-slim

# install curl for health check
RUN apt-get update && \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=go-builder  /code/build/minitiad /usr/local/bin/minitiad

# Setup minitia user
WORKDIR /minitia
RUN addgroup minitia \
    && adduser --ingroup minitia --disabled-password --home /minitia minitia
RUN chown -R minitia:minitia /minitia
USER minitia

# rest server
EXPOSE 1317
# grpc
EXPOSE 9090
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# geth jsonrpc
EXPOSE 8545
# geth jsonrpc ws
EXPOSE 8546

CMD ["/usr/local/bin/minitiad", "version"]
