package feed

import (
	"encoding/xml"
)

// URL is an unparsed URL.
type URL = string

type Channel struct {
	XMLName xml.Name `xml:"channel"`

	// Title of the podcast. Required.
	Title string `xml:"title"`

	// Description of the podcast. Required.
	Description string `xml:"description"`

	// ItunesImage is for the podcast artwork. Required.
	ItunesImage ItunesImage

	// Language of the podcast. Required.
	Language string `xml:"language"`

	// ItunesCategories of the podcast. Required.
	ItunesCategories []ItunesCategory

	// ItunesExplicit marks explicit content. Required.
	ItunesExplicit bool `xml:"itunes:explicit"`

	// Items are the podcast episodes. Required.
	Items []Item
}

// ItunesImage is for an image URL.
type ItunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    URL      `xml:"href,attr"`
}

// ItunesCategory is a category and optional subcategory.
type ItunesCategory struct {
	XMLName           xml.Name           `xml:"itunes:category"`
	Text              string             `xml:"text,attr"`
	ItunesSubCategory *ItunesSubCategory `xml:",omitempty"`
}

// ItunesSubCategory is optional and always sits below ItunesCategory.
type ItunesSubCategory struct {
	XMLName xml.Name `xml:"itunes:category"`
	Text    string   `xml:"text,attr"`
}
