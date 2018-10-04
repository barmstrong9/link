package main

import (
	"fmt"
	"strings"

	"github.com/barmstrong9/link"
)

var exampleHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat thing <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
