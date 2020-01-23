package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	feed "subbu.com/learngo/feed"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.Info("main Init complete")
}

func feedsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Got Request Header : ", r.Header)
	w.Header().Set("Content-Type", "application/json")
	resp := loadFeeds()
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	encoder.Encode(resp)
}

func loadFeeds() map[string]feed.Result {
	var urls []string
	var err error
	urls, err = feed.ReadFeedLinks("/var/tmp/feedlist")

	if err != nil {
		log.Error("error reading feed links : ", err)
	}

	return feed.GetFeeds(urls, 10)
}

func main() {
	http.HandleFunc("/feeds", feedsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
