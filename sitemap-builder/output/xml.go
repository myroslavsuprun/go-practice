package output

import (
	"encoding/xml"
	"fmt"
)

func XmlParse(links map[string]string) (string, error) {
	type Url struct {
		Loc string `xml:"loc"`
	}

	type Urlset struct {
		XMLName xml.Name `xml:"urlset"`
		Xmlns   string   `xml:"xmlns,attr"`
		Urls    []Url    `xml:"url"`
	}

	urls := []Url{}
	for k, _ := range links {
		urls = append(urls, Url{k})
	}

	urlset := Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	out, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error while marshaling xml: %v", err)
	}

	return withHeader(out), nil
}

func withHeader(out []byte) string {
	return xml.Header + string(out)
}
