package feed_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/matryer/is"

	"github.com/maragudk/podcasthost/feed"
)

//go:embed testdata/required.xml
var required string

func TestDocument_Render(t *testing.T) {
	t.Run("renders a feed in xml", func(t *testing.T) {
		is := is.New(t)

		d := feed.New()
		d.Channel.Title = "The Podcast Show"
		d.Channel.Description = "A show about podcasts."
		d.Channel.ItunesImage.Href = "https://example.com/artwork.jpg"
		d.Channel.Language = "en-us"
		d.Channel.ItunesCategories = append(d.Channel.ItunesCategories, feed.ItunesCategory{
			Text: "Education",
			ItunesSubCategory: &feed.ItunesSubCategory{
				Text: "How To",
			},
		})
		d.Channel.ItunesExplicit = true
		d.Channel.Items = []feed.Item{
			{
				Title: "Episode 1",
				Enclosure: feed.Enclosure{
					URL:    "https://example.com/1.mp4",
					Length: "123",
					Type:   "audio/mp4",
				},
			},
		}

		var b bytes.Buffer
		err := d.Render(&b)
		is.NoErr(err)
		is.Equal(b.String(), required)
	})
}
