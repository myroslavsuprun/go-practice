package main

import (
	"golang.org/x/net/html"

	"fmt"
	"os"

	"link/explorer"
	"link/links"
)

func main() {
	fR, err := os.Open("./ex4.html")
	defer fR.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	node, err := html.Parse(fR)
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes := explorer.Get(node, explorer.WithConstraint(func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "a"
	}))

	links := links.ParseLinks(&nodes)

	fmt.Println(links)
}
