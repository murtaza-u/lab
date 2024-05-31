package main

import (
	"fmt"

	"github.com/reujab/wallpaper"
)

func main() {
	background, _ := wallpaper.Get()
	fmt.Println("Current wallpaper:", background)
	wallpaper.SetFromFile("/home/murtaza/Pictures/wallpapers/wallhaven-4d7w7l.jpg")
}
