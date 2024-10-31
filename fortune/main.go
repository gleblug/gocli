package main

/*
 * Inspired by Unix program:
 * https://en.wikipedia.org/wiki/Fortune_(Unix)
 *
*/

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	fortunesPath = "fortunes"
)

type Fortune struct {
	source string
	text string
}

func main() {
	// TODO: add more flags
	fortunes := getFortunes(fortunesPath)
	index := randomInt(len(fortunes))
	fortune := fortunes[index]
	fmt.Println(fortune.text)
}

func getFortunes(root string) []Fortune {
	entries, err := os.ReadDir(root)
	if err != nil {
		log.Fatal("can't get fortunes:", err)
	}

	res := make([]Fortune, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		filename := e.Name()
		content, err := os.ReadFile(filepath.Join(root, filename))
		if err != nil {
			log.Fatal("can't open file: ", err)
		}

		res = append(res, Fortune{
			source: filename,
			text: string(content),
		})
	}

	return res
}

func randomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}