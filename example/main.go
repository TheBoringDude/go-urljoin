package main

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
)

func main() {
	a, _ := urljoin.UrlJoin("http://sample.com", "//hello//", "?test=1", "?hello=s")

	fmt.Println(a)
}
