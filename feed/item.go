package feed

import (
	"encoding/xml"
)

// Item is a podcast episode.
type Item struct {
	XMLName xml.Name `xml:"item"`

	// Title of the episode. Required.
	Title string `xml:"title"`

	// Enclosure holds episode content, file size, and type information. Required.
	Enclosure Enclosure
}

// MIMEType is just a string alias for now, but may provide validation later.
type MIMEType = string

// Enclosure holds episode content, file size, and type information.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     URL      `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    MIMEType `xml:"type,attr"`
}
