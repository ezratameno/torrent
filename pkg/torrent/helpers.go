package torrent

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func CheckIfEpisodeBelongToADesiredShow(c *TorrentClient, episodeName string) (bool, string) {
	for _, showName := range c.ShowsName {
		if strings.Contains(strings.ToLower(episodeName), strings.ToLower(showName)) {
			return true, showName
		}
	}
	return false, ""
}

// get the specific episode data like download link
func GetEpisodeData(episode *EpisodeData) {
	// episodeNum := strings.Split(strings.TrimPrefix(episode.Name, episode.ShowName), ".")
	// if len(episodeNum) > 0 {
	// 	episode.EpisodeNum = episodeNum[1]
	// }
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)
	c.OnHTML(".dropdown", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, h *colly.HTMLElement) {
			if h.Text == "ITORRENTS MIRROR" {
				episode.DownloadLink = h.Attr("href")
			}
		})
	})
	c.OnHTML(".list", func(e *colly.HTMLElement) {
		var varName string
		e.ForEach("li", func(i int, li *colly.HTMLElement) {
			li.ForEach("strong", func(index int, e *colly.HTMLElement) {
				varName = e.Text
			})
			li.ForEach("span", func(i int, span *colly.HTMLElement) {
				switch varName {
				case "Category":
					episode.Category = span.Text
				case "Type":
					episode.Type = span.Text
				case "Language":
					episode.Language = span.Text
				case "Total size":
					episode.TotalSize = span.Text
				case "Uploaded By":
					episode.Uploader = span.Text
				case "Downloads":
					episode.Downloads = span.Text
				case "Seeders":
					episode.Seeders = span.Text
				case "Leechers":
					episode.Leechers = span.Text
				}

			})

		})
	})
	var imgCounter int = 0
	c.OnHTML(".img-responsive", func(e *colly.HTMLElement) {
		// make sure we only take the first photo
		if imgCounter == 0 {
			episode.ImageURL = e.Attr("data-original")
			imgCounter++
		}
	})

	c.Visit(episode.Link)

}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
