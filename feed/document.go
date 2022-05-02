package feed

import (
	"encoding/xml"
	"io"
)

const (
	ItunesNamespace  = "http://www.itunes.com/dtds/podcast-1.0.dtd"
	ContentNamespace = "http://purl.org/rss/1.0/modules/content/"
)

// Document is the top-level XML document named <rss>.
type Document struct {
	XMLName          xml.Name `xml:"rss"`
	Version          string   `xml:"version,attr"`
	ItunesNamespace  string   `xml:"xmlns:itunes,attr"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Channel          Channel
}

// New Document with valid version and namespaces.
func New() Document {
	return Document{
		Version:          "2.0",
		ItunesNamespace:  ItunesNamespace,
		ContentNamespace: ContentNamespace,
	}
}

func (d Document) Render(w io.Writer) error {
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(d); err != nil {
		return err
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
