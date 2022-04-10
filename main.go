package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type randomFact struct {
	Text string `json:"test"`
}

func getRandomFacts() (randomFact, error) {
	var fact randomFact
	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return randomFact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return randomFact{}, err
	}

	return fact, nil
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("Get Useless Fact")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Get Useless Fact", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	win.SetContent(widget.NewLabel("Hello World"))
	win.ShowAndRun()
}
