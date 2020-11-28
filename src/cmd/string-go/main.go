package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

// matches against anything that looks like a delimited string literal. It
// balance the prefix and suffix to be the same delimiter, because the scanner
// will do that for us.
var delimitedPattern = regexp.MustCompile(`(?s)^\^+` + "`" + `(.*)` + "`" + `\^+$`)

func main() {
	flag.Parse()
	fset := token.NewFileSet()
	file := flag.Arg(0)
	f, err := parser.ParseFile(fset, file, nil, 0)
	check(err)
	// Rewrite source to quoted literal
	astutil.Apply(f, func(c *astutil.Cursor) bool {
		switch n := c.Node(); n := n.(type) {
		case *ast.BasicLit:
			if n.Kind != token.STRING {
				break
			}
			matches := delimitedPattern.FindStringSubmatch(n.Value)
			if len(matches) != 2 {
				break
			}
			n.Value = strconv.Quote(matches[1])
			c.Replace(n)
		}
		return true
	}, nil)
	buf := new(bytes.Buffer)
	check(format.Node(buf, fset, f))
	outfile := filepath.Join(filepath.Dir(file), "generated."+filepath.Base(file))
	check(ioutil.WriteFile(outfile, buf.Bytes(), 0666))
}
