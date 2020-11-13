# string-go

This is a fork of Go at [golang/go@30ba798](https://github.com/golang/go/commit/30ba7980932dfb7ec6660ee929b4e1982256285f) to add
delimited raw string literals to the language. See [golang/go#32590](https://github.com/golang/go/issues/32590).

To try it out, execute
```
$ go run cmd/string-go file.go
```
where `file.go` contains Go source with delimited strings. Running this command
will place `generated.file.go` into the current directory, containing the equivalent
Go source without delimited strings.
