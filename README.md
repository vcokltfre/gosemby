# gosemby

A VM for [Semby](https://github.com/vcokltfre/semby) written in Go.

## Using gosemby

To use gosemby you need to build the project with Go:

```sh
go build
```

You can now put the `gosemby` binary somewhere in your path, for example `/usr/bin`.

To run a bytecode file with gosemby you can run:

```sh
gosemby <file>
```

To run a Semby source file with gosemby you can run `semby` with the vm flag:

```sh
semby run <file> --vm gosemby
```
