package stories

import (
	"encoding/json"
	"io"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func Get() (Story, error) {
	story, err := getStoryFile()
	if err != nil {
		return nil, err
	}
	return story, nil
}

func getStoryFile() (Story, error) {
	data, err := readStoryFile()
	if err != nil {
		return nil, err
	}

	story, err := parseStoryFile(data)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func readStoryFile() ([]byte, error) {
	file, err := os.Open("./gopher.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseStoryFile(data []byte) (Story, error) {
	var story Story
	err := json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
