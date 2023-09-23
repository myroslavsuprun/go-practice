package templates

import (
	"cyoa/stories"
	"html/template"
)

func Get() (*template.Template, error) {
	t, err := createTmpl()
	if err != nil {
		return nil, err
	}

	return t, nil

}

func createTmpl() (*template.Template, error) {
	tmpl, err := template.New("index.html").ParseFiles("./assets/index.html")
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

type tmplData struct {
	Title      string
	Paragraphs []string
	Options    []option
}

type option struct {
	Text    string
	Chapter string
}

func ParseTmplChapter(ch stories.Chapter) tmplData {
	return tmplData{
		Title:      ch.Title,
		Paragraphs: ch.Paragraphs,
		Options:    parseTmplOptions(ch.Options),
	}
}

func parseTmplOptions(opts []stories.Option) []option {
	var tmplOpts []option
	for _, opt := range opts {
		tmplOpts = append(tmplOpts, option{
			Text:    opt.Text,
			Chapter: opt.Chapter,
		})
	}
	return tmplOpts
}
