package main

import (
	"ezratameno/torrent/pkg/torrent"
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	client := torrent.NewTorrentClient()
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("hello")
	// fix starting size
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
			object.(*widget.Label).Text = client.TodayEpisodes[id].ShowName
		})

	// right side container
	contentText := widget.NewLabel("Please select a show")
	contentText.Wrapping = fyne.TextWrapWord
	rightContainer := container.NewMax(contentText)
	rightContainer.Add(contentText)
	cur, _ := os.Getwd()
	fmt.Println(cur)
	listView.OnSelected = func(id widget.ListItemID) {
		contentText.Text = formatData(client.TodayEpisodes[id])
		btn := widget.NewButton("Download", func() {})
		btn.OnTapped = func() {
			client.Download(client.TodayEpisodes[id])
		}
		// img := canvas.NewImageFromFile("C:/Users/ezrat/Desktop/Go_projects/torrent/pkg/torrent/images/bg.png")
		rightContainer.Add(btn)
		// rightContainer.Add(img)

	}

	split := container.NewHSplit(
		listView,
		rightContainer,
	)
	split.Offset = 0.2
	w.SetContent(split)
	w.ShowAndRun()
}

func formatData(episode torrent.EpisodeData) string {
	return fmt.Sprintf("Show Name:	%s\nEpisode Name:	%s\nEpisode Number:		%s\nEpisode Size:	%s\n",
		episode.ShowName, episode.Name, episode.EpisodeNum, episode.TotalSize)
}
