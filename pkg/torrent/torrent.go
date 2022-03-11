package torrent

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gocolly/colly"
)

var listOfShows []string = []string{"Young.Justice", "Legacies", "Star.Trek.Discovery", "bmf.", "power.book.ii.ghost", "S.W.A.T", "Blue.Bloods", "SEAL.Team", "Dexter", "Peaky.Blinders"}

const domain string = "1337x.to"
const fileType string = ".torrent"

type EpisodeData struct {
	Name         string
	EpisodeNum   string
	ShowName     string
	Link         string
	DownloadLink string
	Category     string
	Type         string
	Language     string
	TotalSize    string
	Uploader     string
	Downloads    string
	Seeders      string
	Leechers     string
	ImageURL     string
}
type TorrentClient struct {
	ShowsName     []string
	TodayEpisodes []EpisodeData
}

func NewTorrentClient() *TorrentClient {
	client := &TorrentClient{
		ShowsName: listOfShows}
	client.TodayEpisodes = client.getTodayEpisodes()
	return client
}

func (client *TorrentClient) getTodayEpisodes() []EpisodeData {
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	episodesFromScarpper := []EpisodeData{}
	c.OnHTML("td", func(e *colly.HTMLElement) {
		// the name of the episode and the link to the page are always on index 1
		// of each row in the table
		var index int
		e.ForEach("a", func(i int, h *colly.HTMLElement) {
			if index == 1 {
				ep := EpisodeData{Name: e.Text, Link: "https://" + domain + h.Attr("href")}
				episodesFromScarpper = append(episodesFromScarpper, ep)
				index = 0
				return
			}
			index++
		})
	})
	// Start scraping on https://1337x.to/popular-tv
	c.Visit("https://1337x.to/popular-tv")
	c.Visit("https://1337x.to/popular-movies")

	return GetEpisodesData(episodesFromScarpper)
}

func GetEpisodesData(episodes []EpisodeData) []EpisodeData {
	var res []EpisodeData
	for _, episodeInstance := range episodes {
		GetEpisodeData(&episodeInstance)
		res = append(res, episodeInstance)
	}
	return res
}
func (c *TorrentClient) AddDesiredShow(show string) {
	c.ShowsName = append(c.ShowsName, show)
	c.TodayEpisodes = c.getTodayEpisodes()
}

func (client *TorrentClient) Download(episode EpisodeData) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := home + "/Downloads/"
	if runtime.GOOS == "windows" {
		filePath = filepath.FromSlash(filePath)
	}
	err = DownloadFile(filePath+episode.Name+fileType, episode.DownloadLink)
	if err != nil {
		return nil
	}
	err = client.openTorrentFile(episode)
	if err != nil {
		return err
	}
	return nil
}

func (client *TorrentClient) openTorrentFile(episode EpisodeData) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := home + "/Downloads/"
	cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", filePath+episode.Name+fileType)
	cmd.Start()
	return nil
}
