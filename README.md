# MelonWallet Gateway Service

---------------------
The MelonWallet microservice responsible for service routing, rate-limiting, and account authentication.

## Compile and build

---------------------

```bash
bazel build //...
```

## Initialize a BUILD for a new folder

---------------------

```bash
bazel run //:gazelle
```

## Update dependencies for Bazel build

---------------------

```bash
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

## Run tests

---------------------

```bash
bazel test //...
```
