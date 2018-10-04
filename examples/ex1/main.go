package main

import (
	"fmt"
	"strings"

	"github.com/barmstrong9/link"
)

var exampleHTML = `
<html>
	<body>
  		<h1>Hello!</h1>
		  <a href="/other-page">
		  A link to another page
		  <span>It's just a span bro<span>
		  </a>

		  <a href="/another_one">A DJ Khaled Link</a>
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
