package main

import (
	"ezratameno/torrent/pkg/torrent"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func printListOfFoundDesiredEpisodes(episodes []torrent.EpisodeData) {
	colorGreen := "\033[32m"
	colorWhite := "\033[37m"

	fmt.Println(string(colorGreen), "Episodes List:")
	fmt.Println(string(colorWhite))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Show Name", "Episode Number", "Total Size"})
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	var data [][]string
	var counter int
	for _, episode := range episodes {
		tmp := []string{strconv.Itoa(counter), episode.ShowName, episode.EpisodeNum, episode.TotalSize}
		data = append(data, tmp)
		counter++
	}
	table.AppendBulk(data)
	table.Render()
}

func showTodayEpisodes(client *torrent.TorrentClient) {
	if len(client.TodayEpisodes) > 0 {
		printListOfFoundDesiredEpisodes(client.TodayEpisodes)
	} else {
		fmt.Println("Didn't found any shows today :(")
	}
}

func displayShows(client *torrent.TorrentClient) {
	fmt.Println("The list of shows that i am intresed in are:")
	var counter = 1
	for _, name := range client.ShowsName {
		msg := fmt.Sprintf("%d) %s", counter, name)
		fmt.Println(msg)
		counter++
	}
}

func clearScreen() {
	fmt.Println("Press Enter to go back to the main menu...")
	fmt.Scanln()
	// need to make cross platform
	// clear screen
	command := []string{"/c", "cls"}
	cmd := exec.Command("cmd", command...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
