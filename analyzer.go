package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func InitialAction(f string, action bool) int {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//scan the file line by line, stop at the first line
	scanner := bufio.NewScanner(file)
	var firstline string
	for scanner.Scan() {
		firstline = scanner.Text()
		break
	}
	//scan the line word by word.
	//if action file then the timestamp is the second token
	//if episode file the the timestamp is the first token
	wScanner := bufio.NewScanner(strings.NewReader(firstline))
	wScanner.Split(bufio.ScanWords)
	var initialTs int
	c := 0
	for wScanner.Scan() {
		if action {
			if c == 1 {
				initialTs, err = strconv.Atoi(wScanner.Text())
				if err != nil {
					panic(err)
				}
				break
			}
			c++
		}
		if !action {
			initialTs, err = strconv.Atoi(wScanner.Text())
			if err != nil {
				panic(err)
			}
			break
		}

	}
	return initialTs
}

func FirstEpisodeDuration(initial int) int {
	firstEpisode := InitialAction("besouroEpisodes.txt", false)
	return firstEpisode - initial
}
