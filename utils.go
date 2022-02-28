package main

import (
	"bufio"
	"math/rand"
	"os"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func readWordlist(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var wl []string
	for scanner.Scan() {
		wl = append(wl, scanner.Text())
	}

	return wl, nil
}
