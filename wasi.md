# How to compile WASI executable

```shell
$ GOOS=wasip1 GOARCH=wasm go build -o main.wasm
```

# How to run WASI executable

## wasmtime

```shell
$ wasmtime main.wasm
```

## wasmedge

```shell
$ wasmedge main.wasm
```

