package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#fu678",
	}

	colors["white"] = "#js7654"

	fmt.Println(colors)
}
