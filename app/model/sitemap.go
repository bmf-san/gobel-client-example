package model

// A URLSet represents the singular of urlset.
type URLSet struct {
	XMLName string `xml:"urlset"`
	URLs    []URL  `xml:"url"`
}

// URL represents the singular of url.
type URL struct {
	Loc string `xml:"loc"`
}
