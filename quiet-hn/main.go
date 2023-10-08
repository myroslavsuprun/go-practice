package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"
)

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type TempStory struct {
	Title string
	URL   string
	Host  string
}

type TempData struct {
	Stories []TempStory
	Time    string
}

type Cache struct {
	Stories []TempStory
	Time    time.Time
}

func main() {
	limit := flag.Int("stories", 30, "Number of stories to fetch")
	flag.Parse()

	temp := template.Must(template.ParseFiles("news.html"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHandler(limit, temp))

	fmt.Println("Server is starting on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func getHandler(limit *int, temp *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ts := time.Now()

		st, err := getStories(limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}

		tf := time.Now()
		timeTaken := tf.Sub(ts)
		timeTaken = timeTaken.Round(time.Millisecond)
		tempData := TempData{
			Stories: st,
			Time:    timeTaken.String(),
		}

		err = temp.Execute(w, tempData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}
}

var (
	storiesCache Cache
	muxCache     sync.Mutex
)

func getStories(limit *int) ([]TempStory, error) {
	muxCache.Lock()
	defer muxCache.Unlock()
	if storiesCache.Stories != nil && time.Now().Sub(storiesCache.Time) < time.Minute*10 {
		return storiesCache.Stories, nil
	}

	topStories, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, err
	}

	var stories []int
	err = json.NewDecoder(topStories.Body).Decode(&stories)
	if err != nil {
		return nil, err
	}

	stories = stories[:*limit]
	type channelledStory struct {
		idx   int
		story Story
	}
	var storyChan = make(chan channelledStory, len(stories))

	for idx, id := range stories {
		id := id
		idx := idx
		go func() {
			story, _ := getStory(id)
			storyChan <- channelledStory{
				idx:   idx,
				story: story,
			}
		}()
	}

	tStories := make([]TempStory, len(stories))
	for i := 0; i < len(stories); i++ {
		chData := <-storyChan
		tStories[chData.idx] = TempStory{
			Title: chData.story.Title,
			URL:   chData.story.URL,
			Host:  chData.story.URL,
		}
	}

	storiesCache = Cache{
		Stories: tStories,
		Time:    time.Now(),
	}
	return tStories, nil
}

func getStory(id int) (Story, error) {
	var story Story
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + fmt.Sprint(id) + ".json")
	if err != nil {
		return story, err
	}

	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		return story, err
	}

	return story, nil
}
