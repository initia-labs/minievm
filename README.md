# MiniEVM

MiniEVM is an optimistic rollup consumer chain powered by EVM, designed to simplify the process of bootstrapping an L2 network. The main advantage of using MiniEVM is that the users can leverage the OPinit stack for enhanced security and utilize all the Initia ecosystem tooling from day one, without the need to prepare a validator group or build the users' own ecosystem tools.

- [go-ethereum](https://github.com/initia-labs/evm)

## Prerequisites

- Go v1.23.3+
- (optional) [solc-select](https://github.com/crytic/solc-select) v1.0.4+ (used in contractsgen.sh)

## Getting Started

To get started with L2, please visit the [documentation](https://docs.initia.xyz).

## Features

- Powered by EVM, MiniEVM acts as an optimistic rollup consumer chain.
- Simplifies the network bootstrapping process, making it faster and more efficient.
- Eliminates the need for setting up a validator group or creating custom ecosystem tools.
- Integrates seamlessly with the OPinit stack, enhancing security.
- Provides immediate access to the full suite of Initia ecosystem tools right from the start.

## JSON-RPC

For information of supported JSON-RPC methods, please refer to the [JSON-RPC documentation](jsonrpc/README.md).

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the [BSL License](LICENSE).
