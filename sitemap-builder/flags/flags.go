package flags

import (
	"flag"
	"fmt"
)

type Input struct {
	Url   string
	Depth int
}

func Parse() (Input, error) {
	var url string
	depth := 3
	flag.StringVar(&url, "url", "", "url to parse")
	flag.IntVar(&depth, "depth", depth, "depth of urls to parse (default 3)")
	flag.Parse()

	if url == "" {
		return Input{}, fmt.Errorf("url flag is required")
	}

	return Input{
		Url:   url,
		Depth: depth,
	}, nil
}
