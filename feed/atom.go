package feed

const (
	selfLink = "self"
)

// Atom feed format
type Atom struct {
	ID      string `xml:"id" json:"id"`
	Title   string `xml:"title" json:"title"`
	Updated string `xml:"updated" json:"updated"`
	Author  struct {
		Name  string `xml:"name" json:"name"`
		Email string `xml:"email" json:"email"`
		URI   string `xml:"uri" json:"uri"`
	} `xml:"author" json:"author"`
	Links []struct {
		Rel  string `xml:"rel,attr" json:"rel"`
		Href string `xml:"href,attr" json:"href"`
	} `xml:"link" json:"link"`
	Contributor struct {
		Name  string `xml:"name" json:"name"`
		Email string `xml:"email" json:"email"`
		URI   string `xml:"uri" json:"uri"`
	} `xml:"contributor" json:"contributor"`
	Generator string  `xml:"generator" json:"generator"`
	Icon      string  `xml:"icon" json:"icon"`
	Logo      string  `xml:"logo" json:"logo"`
	Rights    string  `xml:"rights" json:"rights"`
	Subtitle  string  `xml:"subtitle" json:"subtitle"`
	Entries   []Entry `xml:"entry" json:"entry"`
}

// Entry in a atom feed
type Entry struct {
	ID      string `xml:"id" json:"id"`
	Title   string `xml:"title" json:"title"`
	Updated string `xml:"updated" json:"updated"`
	Author  struct {
		Name  string `xml:"name" json:"name"`
		Email string `xml:"email" json:"email"`
		URI   string `xml:"uri" json:"uri"`
	} `xml:"author" json:"author"`
	Link struct {
		Href string `xml:"href,attr" json:"href"`
	} `xml:"link" json:"link"`
	Content  string `xml:"content" json:"content"`
	Summary  string `xml:"summary" json:"summary"`
	Category struct {
		Term string `xml:"term,attr" json:"term"`
	} `xml:"category" json:"category"`
	Contributor struct {
		Name  string `xml:"name" json:"name"`
		Email string `xml:"email" json:"email"`
		URI   string `xml:"uri" json:"uri"`
	} `xml:"contributor" json:"contributor"`
	Published string `xml:"published" json:"published"`
}
