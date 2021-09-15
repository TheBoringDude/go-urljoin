package main

import (
	"fmt"

	urljoin "github.com/TheBoringDude/go-urljoin"
)

func main() {
	a := urljoin.UrlJoin("sample.com")

	fmt.Println(a)
}
