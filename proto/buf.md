# Protobufs

This is the public protocol buffers API for [MiniEVM](https://github.com/initia-labs/minievm).

## npm Package

TypeScript definitions are published to npm as [`@initia/minievm-proto`](https://www.npmjs.com/package/@initia/minievm-proto) on every tagged release (`v*`).

### Installation

```bash
npm install @initia/minievm-proto @bufbuild/protobuf
```

### Usage

```typescript
import { MsgCreateSchema } from "@initia/minievm-proto/minievm/evm/v1/tx_pb.js";
import { ParamsSchema } from "@initia/minievm-proto/minievm/evm/v1/types_pb.js";
```

The package requires `@bufbuild/protobuf` v2 as a peer dependency.
