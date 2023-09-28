package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"link"

	"sitemap-builder/flags"
	"sitemap-builder/output"
)

func main() {
	i, err := flags.Parse()
	if err != nil {
		log.Fatal(err)
	}

	baseUrl := getBaseUrl(i.Url)
	totalLinks := make(map[string]string)
	urlsToParse := []string{i.Url}
	count := 0

	for len(urlsToParse) > 0 && count < i.Depth {
		url := urlsToParse[0]
		urlsToParse = urlsToParse[1:]

		links, err := getLinksByUrl(url)
		if err != nil {
			continue
		}

		for _, l := range links {
			href := addBaseUrl(l.Href, baseUrl)

			if _, ok := totalLinks[href]; ok {
				continue
			}

			totalLinks[href] = l.Text

			if strings.HasPrefix(href, baseUrl) {
				urlsToParse = append(urlsToParse, href)
			}
		}

		count++
	}

	out, err := output.XmlParse(totalLinks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func addBaseUrl(href string, baseUrl string) string {
	var url string
	if strings.HasPrefix(href, "http") {
		url = href
	} else if strings.HasPrefix(href, "/") {
		url = fmt.Sprintf("%s%s", baseUrl, href)
	} else {
		url = fmt.Sprintf("%s/%s", baseUrl, href)
	}

	return url
}

func getBaseUrl(url string) string {
	split := strings.Split(url, "/")
	return fmt.Sprintf("%s//%s", split[0], split[2])
}

func getLinksByUrl(url string) ([]link.Link, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error getting URL: %v", err)
	}
	defer resp.Body.Close()

	links, err := getLinks(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing links: %v", err)
	}

	return links, nil
}

func getLinks(r io.Reader) ([]link.Link, error) {
	links, err := link.Parse(&r)
	if err != nil {
		return nil, err
	}

	return links, nil
}
