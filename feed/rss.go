package feed

// Rss core rss feed
type Rss struct {
	Title       string `xml:"channel>title" json:"title"`
	Link        string `xml:"channel>link" json:"link"`
	Description string `xml:"channel>description" json:"description"`
	Image       struct {
		URL         string `xml:"url" json:"url"`
		Title       string `xml:"title" json:"title"`
		Link        string `xml:"link" json:"link"`
		Width       string `xml:"width" json:"width"`
		Height      string `xml:"height" json:"height"`
		Description string `xml:"description" json:"description"`
	} `xml:"channel>image" json:"image"`
	Language      string `xml:"channel>language"  json:"language"`
	Copyright     string `xml:"channel>copyright"  json:"copyright"`
	PubDate       string `xml:"channel>pubDate"  json:"pubDate"`
	LastBuildDate string `xml:"channel>lastBuildDate"  json:"lastBuildDate"`
	Category      string `xml:"channel>category"  json:"category"`
	Generator     string `xml:"channel>generator"  json:"generator"`
	TTL           string `xml:"channel>ttl"  json:"ttl"`
	RssItems      []Item `xml:"channel>item"  json:"item"`
}

// Item is a single item in the RSS feed
type Item struct {
	Title       string `xml:"title" json:"title"`
	Link        string `xml:"link" json:"link"`
	Description string `xml:"description" json:"description"`
	Author      string `xml:"author" json:"author"`
	Category    string `xml:"category" json:"category"`
	Comments    string `xml:"comments" json:"comments"`
	Enclosure   struct {
		URL    string `xml:"url,attr" json:"url,attr"`
		Length string `xml:"length,attr" json:"length,attr"`
		Type   string `xml:"type,attr" json:"type,attr"`
	} `xml:"enclosure" json:"enclosure"`
	GUUID   string `xml:"guid" json:"guid"`
	PubDate string `xml:"pubDate" json:"pubDate"`
	Source  string `xml:"source" json:"source"`
}
