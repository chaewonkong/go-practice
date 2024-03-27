package main

import (
	"fmt"
	"os"
)

// practice: using channel

func main() {
	content := make(chan string)

	go func() {
		c, err := os.ReadFile("example.txt")
		if err != nil {
			close(content)
			return
		}
		content <- string(c)
	}()

	contents, ok := <-content
	if !ok {
		fmt.Println("Error reading file")
		return
	}

	fmt.Println(contents)
	close(content)
}
