package model

// A Meta represents the singular of meta.
type Meta struct {
	Title         string
	Canonical     string
	Description   string
	OGTitle       string
	OGDescription string
	OGURL         string
	OGType        string
	OGImage       string
	OGSiteName    string
	OGLocale      string
	TwitterCard   string
	TwitterSite   string
	NoIndex       bool
}
