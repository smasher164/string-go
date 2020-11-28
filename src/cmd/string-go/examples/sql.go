package a

const s = ^^`
func main() {
	sql := ^`SELECT `foo` FROM `bar` WHERE `baz` = "qux"`^
	fmt.Println(sql)
}
`^^