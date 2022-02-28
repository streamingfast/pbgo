# Change log

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Changed

- **BREAKING CHANGE** Function `github.com/streamingfast/pbog/sf/blockmeta/v1/pbblockmeta#NewClient` now accepts a `conn *grpc.Conn` directly instead of an `addr string` value. This has been made to remove the cycle we had on `dgrpc` which was depending on `pbgo` library. It also does not return an `error` anymore. To get back the old behavior, change:

  ```golang
  client, err := pbblockmeta.NewClient(blockmetaAddr)
  if err != nil {
      // Error handling
  }
  ```

  to

  ```golang
  conn, err := dgrpc.NewInternalClient(blockmetaAddr)
  if err != nil {
      // Error handling
  }

  client := pbblockmeta.NewClient(conn)
  ```

- Re-generated everything with `google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1` and `google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0`, this means this library now requires a recent enough `google.golang.org/grpc` library (`1.44` listed here).

### Removed

- **BREAKING CHANGE** Removed `grpc.health.v1.health` generated service, if you were depending on our generated code, replace `github.com/streamingfast/pbgo/grpc/health/v1` by `google.golang.org/grpc/health/grpc_health_v1` in your Golang imports.

- Removed dependency on `github.com/golang/protobuf`, replaced by the modern replacement `google.golang.org/protobuf`.

## 2020-03-21

### Changed

* License changed to Apache 2.0