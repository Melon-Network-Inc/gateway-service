# MelonWallet Gateway Service

[![Build status](https://badge.buildkite.com/55efa927385f2f57f050295d5f2416ad0e67db5051ed68372f.svg)](https://buildkite.com/melon-network-inc/gateway-service-pipeline)

<img src="https://avatars.githubusercontent.com/u/104064333?s=400&u=fe08053ed0a72719e2ea4bb0229766ef9b4fdfee&v=4" width="100">

---------------------

The MelonWallet microservice responsible for service routing, rate-limiting, and account authentication.

## Project Setup

### Compile and build

```bash
bazel build //...
```

### Start payment server

```bash
bazel run cmd/server:server
```

### Update dependencies for Bazel build

```bash
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

### Run tests

```bash
bazel test //...
```

### Update API document

```bash
python3 ../NestedJsonMerger/merge.py
```

## Swagger Doc

Staging URL
<http://34.168.151.218:8080/swagger/index.html>

Development URL
<http://localhost:8080/swagger/index.html>
