package explorer

import "golang.org/x/net/html"

type executionType int

const (
	preExecution executionType = iota
	postExecution
	nSiblExecution
)

type execution = struct {
	exec     func(*html.Node)
	execType executionType
}

func New(exec ...execution) func(*html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		preExecute(n, exec...)

		if n.FirstChild != nil {
			f(n.FirstChild)
		}

		if n.NextSibling != nil {
			nSiblExecute(n, exec...)
			f(n.NextSibling)
		}

		postExecute(n, exec...)
	}

	return f
}

func preExecute(n *html.Node, exec ...execution) {
	for _, ex := range preExecutions(exec...) {
		ex.exec(n)
	}
}

func postExecute(n *html.Node, exec ...execution) {
	for _, ex := range postExecutions(exec...) {
		ex.exec(n)
	}
}

func nSiblExecute(n *html.Node, exec ...execution) {
	for _, ex := range nextSiblingExecutions(exec...) {
		ex.exec(n)
	}
}

func PreExec(f func(*html.Node)) execution {
	return execution{
		execType: preExecution,
		exec: func(n *html.Node) {
			f(n)
		},
	}
}

func PostExec(f func(*html.Node)) execution {
	return execution{
		execType: postExecution,
		exec: func(n *html.Node) {
			f(n)
		},
	}
}

func NextSiblingExecution(f func(*html.Node)) execution {
	return execution{
		execType: nSiblExecution,
		exec: func(n *html.Node) {
			f(n)
		},
	}
}

func nextSiblingExecutions(exec ...execution) []execution {
	var nextSiblingExec []execution
	for _, ex := range exec {
		if ex.execType == nSiblExecution {
			nextSiblingExec = append(nextSiblingExec, ex)
		}
	}
	return nextSiblingExec
}

func preExecutions(exec ...execution) []execution {
	var preExec []execution
	for _, ex := range exec {
		if ex.execType == preExecution {
			preExec = append(preExec, ex)
		}
	}
	return preExec
}

func postExecutions(exec ...execution) []execution {
	var postExec []execution
	for _, ex := range exec {
		if ex.execType == postExecution {
			postExec = append(postExec, ex)
		}
	}
	return postExec
}

// --- First kind of solution
// insideAnchor := false
// exp := explorer.New(
// 	explorer.NexSiblingExecution(func(n *html.Node) {
// 		parent := n.Parent
// 		if parent != nil && parent.Type == html.ElementNode && parent.Data == "a" {
// 			insideAnchor = false
// 		}
// 	}),
// 	explorer.PreExecution(
// 		func(n *html.Node) {
// 			if n.Type == html.ElementNode && n.Data == "a" {
// 				insideAnchor = true
//
// 				for _, a := range n.Attr {
// 					if a.Key == "href" {
// 						link := Link{Href: a.Val}
// 						links = append(links, link)
// 					}
// 				}
// 			}
// 		},
// 	),
// 	explorer.PreExecution(func(n *html.Node) {
// 		if n.Type == html.TextNode && insideAnchor {
// 			trimmed := n.Data
// 			if len(trimmed) > 0 {
// 				last := &links[len(links)-1]
//
// 				if len(last.Text) > 0 {
// 					last.Text = fmt.Sprintf("%s%s", last.Text, trimmed)
// 				}
//
// 				if len(last.Text) == 0 {
// 					last.Text = trimmed
// 				}
// 			}
// 		}
// 	}),
// )
//
// exp(node)
//

// --- Second kind of solution
// var f func(*html.Node, bool)
// 	f = func(n *html.Node, insideAnchor bool) {
//
// 		if n.NextSibling != nil {
// 			f(n.NextSibling, insideAnchor)
// 		}
//
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			insideAnchor = true
//
// 			for _, a := range n.Attr {
// 				if a.Key == "href" {
// 					link := Link{Href: a.Val}
// 					links = append(links, link)
// 				}
// 			}
// 		}
//
// 		if n.Type == html.TextNode && insideAnchor {
// 			trimmed := n.Data
// 			if len(trimmed) > 0 {
// 				last := &links[len(links)-1]
//
// 				if len(last.Text) > 0 {
// 					last.Text = fmt.Sprintf("%s%s", last.Text, trimmed)
// 				}
//
// 				if len(last.Text) == 0 {
// 					last.Text = trimmed
// 				}
// 			}
// 		}
//
// 		if n.FirstChild != nil {
// 			f(n.FirstChild, insideAnchor)
// 		}
// 	}
//
// 	f(node, false)
//
//
