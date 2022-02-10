package main

import (
	"bufio"
	"ezratameno/torrent/pkg/torrent"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	client := torrent.NewTorrentClient()
	err := menu(client)
	if err != nil {
		fmt.Println(err)
	}
	// client.AddDesiredShow("Peacemaker.2022")
	// displayShows(client)
}

func menu(client *torrent.TorrentClient) error {
	loop := true
	for loop {
		fmt.Println("Welcome to Ezra's torrent app !!")
		fmt.Println("1)Show todays episodes")
		fmt.Println("2)Display my shows")
		fmt.Println("3)Exit")
		fmt.Println("Please enter your choise:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		//TODO: validate user choise
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		switch choice {
		case 1:
			showTodayEpisodes(client)
			clearScreen()
		case 2:
			displayShows(client)
			clearScreen()
		case 3:
			fmt.Println("exiting...")
			time.Sleep(6 * time.Second)
			loop = false
		}

	}
	return nil
}
