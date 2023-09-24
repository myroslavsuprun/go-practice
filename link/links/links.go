package links

import (
	"fmt"
	"link/explorer"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(nodes *[]*html.Node) []Link {
	var links []Link

	explore := explorer.New(explorer.PreExec(func(n *html.Node) {
		if n.Type == html.TextNode {
			parseText(n, &links)
		}
	}))

	for _, linkNode := range *nodes {
		link := Link{}
		for _, attr := range linkNode.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val
			}
		}

		links = append(links, link)
		explore(linkNode.FirstChild)
	}

	trimLinks(&links)

	return links
}

func trimLinks(links *[]Link) {
	for i, link := range *links {
		(*links)[i].Text = strings.TrimSpace(link.Text)
	}
}

func parseText(n *html.Node, links *[]Link) {
	trimmed := strings.TrimSpace(n.Data)
	if len(trimmed) > 0 {
		last := &(*links)[len(*links)-1]

		if len(last.Text) > 0 {
			last.Text = fmt.Sprintf("%s%s", last.Text, trimmed)
		}

		if len(last.Text) == 0 {
			last.Text = trimmed
		}
	}
}
