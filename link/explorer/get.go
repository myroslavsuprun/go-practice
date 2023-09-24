package explorer

import (
	"golang.org/x/net/html"
)

type Constraint func(*html.Node) bool

func Get(n *html.Node, constraints ...Constraint) []*html.Node {
	var nodes []*html.Node

	explore(n, &nodes, &constraints)

	return nodes
}

func WithConstraint(constraint Constraint) func(*html.Node) bool {
	return func(n *html.Node) bool {
		return constraint(n)
	}
}

func explore(n *html.Node, nodes *[]*html.Node, constraints *[]Constraint) {
	if n.FirstChild != nil {
		explore(n.FirstChild, nodes, constraints)
	}

	if n.NextSibling != nil {
		explore(n.NextSibling, nodes, constraints)
	}

	if !areConstraintsMet(n, constraints) {
		return
	}

	*nodes = append(*nodes, n)
}

func areConstraintsMet(n *html.Node, constraints *[]Constraint) bool {
	for _, constraint := range *constraints {
		if !constraint(n) {
			return false
		}
	}

	return true
}
