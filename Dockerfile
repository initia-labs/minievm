# Stage 1: Build the Go project
FROM golang:1.23-alpine3.19 AS go-builder

# Use build arguments for the target architecture
ARG TARGETARCH
ARG GOARCH

# Install necessary packages
RUN set -eux; apk add --no-cache ca-certificates build-base git cmake

WORKDIR /code
COPY . /code/

# Install mimalloc
RUN git clone --depth 1 https://github.com/microsoft/mimalloc; cd mimalloc; mkdir build; cd build; cmake ..; make -j$(nproc); make install
ENV MIMALLOC_RESERVE_HUGE_OS_PAGES=4

RUN LEDGER_ENABLED=false GOARCH=${GOARCH} LDFLAGS="-linkmode=external -extldflags \"-L/code/mimalloc/build -lmimalloc -Wl,-z,muldefs -static\"" make build

FROM alpine:3.19

RUN addgroup minitia \
    && adduser -G minitia -D -h /minitia minitia

WORKDIR /minitia

COPY --from=go-builder  /code/build/minitiad /usr/local/bin/minitiad

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
