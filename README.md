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

For example, given the following file with a delimited string literal:
```
package a

const s = ^^`
func main() {
	sql := ^`SELECT `foo` FROM `bar` WHERE `baz` = "qux"`^
	fmt.Println(sql)
}
`^^
```
`string-go` will rewrite it into
```
package a

const s = "\nfunc main() {\n\tsql := ^`SELECT `foo` FROM `bar` WHERE `baz` = \"qux\"`^\n\tfmt.Println(sql)\n}\n"

```

View all of the examples in [src/cmd/string-go/examples](src/cmd/string-go/examples).