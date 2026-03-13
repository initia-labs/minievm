# Protobufs

This is the public protocol buffers API for [MiniEVM](https://github.com/initia-labs/minievm).

## npm Package

TypeScript definitions are published to npm as [`@initia/minievm-proto`](https://www.npmjs.com/package/@initia/minievm-proto).

- **Tagged releases** (`v*`) are published as `latest` (e.g. `1.0.0`).
- **Main branch** pushes are published as `canary` (e.g. `0.0.0-canary.<short-sha>`).

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
