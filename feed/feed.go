package feed

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
)

const (
	userAgent     = "User-Agent"
	feedUserAgent = "golang_feed_reader/1.0"
)

// Type represents the typ of feed something is
type Type int

const (
	// RssFeed Type represents a RSS feed
	RssFeed Type = iota
	// AtomFeed Type rpresents a Atom feed
	AtomFeed
	// None represents no format some type of error occured while reading the feed
	None
)

// Result when fetching multiple feeds you need feed results or the error
type Result struct {
	*rawResult
	TimeToFetch time.Duration `json:"fetchDuration"`
	Error       error         `json:"error"`
}

type rawResult struct {
	Type Type  `json:"feedType"`
	Rss  *Rss  `json:"rss"`
	Atom *Atom `json:"atom"`
}

// GetFeeds get feeds in parallel at most parallelFactor HTTP requests can be in flight
// at any given time
func GetFeeds(urls []string, parallelFactor int) map[string]Result {
	var wg sync.WaitGroup
	var width = make(chan int, parallelFactor)
	result := map[string]Result{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			width <- 1
			start := time.Now()
			feed, err := fetchFeed(url)
			end := time.Now()
			result[url] = Result{rawResult: feed, Error: err, TimeToFetch: end.Sub(start)}
			wg.Done()
			<-width
		}(url)
	}

	wg.Wait()
	return result
}

// ReadFeedLinks read feed links from a text file each feed link is expected to be in a diff line
func ReadFeedLinks(linksLocation string) ([]string, error) {
	var contents []byte
	var err error
	if contents, err = ioutil.ReadFile(linksLocation); err != nil {
		return nil, err
	}

	token := func(c rune) bool {
		return c == '\n'
	}
	return strings.FieldsFunc(string(contents), token), nil
}

func fetchFeed(url string) (*rawResult, error) {

	if url == "" {
		return nil, fmt.Errorf("cannot fetch data for empty URL")
	}

	var rawResponse []byte
	var err error
	var result *rawResult
	if rawResponse, err = readRawFeed(url); err != nil {
		return nil, err
	}

	if result, err = unmarshalFeed(rawResponse); err != nil {
		return nil, err
	}
	return result, nil
}

func decode(data []byte, v interface{}) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

func unmarshalFeed(data []byte) (*rawResult, error) {
	rss := &Rss{}
	atom := &Atom{}

	if err := decode(data, rss); err != nil {
		return nil, err
	}
	if rss.Title == "" {
		if err := decode(data, atom); err != nil {
			return nil, err
		}
		return &rawResult{Type: AtomFeed, Rss: nil, Atom: atom}, nil
	}
	return &rawResult{Type: RssFeed, Rss: rss, Atom: nil}, nil
}

func readRawFeed(url string) ([]byte, error) {
	var resp *http.Response
	var err error
	// response body
	var body []byte
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(userAgent, feedUserAgent)

	if resp, err = client.Do(req); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// ok response return body
	if resp.StatusCode == http.StatusOK {
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, fmt.Errorf("received non OK http response %d from %s", resp.StatusCode, url)
}
