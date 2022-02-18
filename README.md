# go-I2P-jpackage

Embeddable, updatable alternative to NSIS based installer for portable Java I2P apps.

The executable produced by this is a standalone binary that can be built for any platform,
but it does not support cross-compilation. It is designed to be `go get` able and `import`
able on any platform supported by Java 17 an Jpackage.

To produce the executable, run:

```sh
go generate
```

instead of go build.

In order to use this *as a library*, you will need Java 14+(17 recommended), bash, Go and
Git installed.
