package main

import (
	"ezratameno/torrent/pkg/torrent"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	client := torrent.NewTorrentClient()
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Ezra's Torrents")
	// fix starting window size
	w.Resize(fyne.NewSize(1200, 800))
	listView := widget.NewList(
		// return the length of the list
		func() int {
			return len(client.TodayEpisodes)
			// create the item that the list is gonna render
		}, func() fyne.CanvasObject {
			return widget.NewLabel("template")
			// update the rendaring of the object
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).Text = client.TodayEpisodes[id].Name
		})
	// right side container
	contentText := widget.NewLabel("Please select a show")
	contentText.Wrapping = fyne.TextWrapWord
	contentContainer := container.NewMax(contentText)
	contentContainer.Add(contentText)
	listView.OnSelected = func(id widget.ListItemID) {
		btn := widget.NewButton("Download", func() {})
		btn.OnTapped = func() {
			client.Download(client.TodayEpisodes[id])
		}

		info := widget.NewLabel(formatData(client.TodayEpisodes[id]))
		// create image
		img, _ := loadResourceFromURLString(client.TodayEpisodes[id].ImageURL)
		image := canvas.NewImageFromResource(img)
		image.FillMode = canvas.ImageFillContain
		infoContainer := container.NewVSplit(image, info)
		// fix initail size
		infoContainer.Offset = 1
		container := container.NewVSplit(infoContainer, btn)
		container.Offset = 1
		contentContainer.Add(container)
	}
	//split the screen into 2
	// left- list
	// right - content of a single item on the list
	split := container.NewHSplit(
		listView,
		contentContainer,
	)
	split.Offset = 0.2
	w.SetContent(split)
	w.ShowAndRun()

}

func formatData(episode torrent.EpisodeData) string {
	return fmt.Sprintf("Episode Name:	%s\nEpisode Number:		%s\nEpisode Size:	%s\nimage:	%s\nLink:	%s",
		episode.Name, episode.EpisodeNum, episode.TotalSize, episode.ImageURL, episode.DownloadLink)
}
