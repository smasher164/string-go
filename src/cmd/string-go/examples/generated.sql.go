package a

const s = "\nfunc main() {\n\tsql := ^`SELECT `foo` FROM `bar` WHERE `baz` = \"qux\"`^\n\tfmt.Println(sql)\n}\n"
