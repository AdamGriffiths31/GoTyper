package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AdamGriffiths31/Typing/database"
	randomtext "github.com/AdamGriffiths31/Typing/randomText"
	"github.com/AdamGriffiths31/Typing/types"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	file, err := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := log.New(file, "Typer: ", log.LstdFlags)
	logger.Println("Typer has started")

	db := database.NewDatabase(logger)
	defer db.BadgerDB.Close()

	mode := flag.String("mode", "", "Sets the mode of the program. Options include: random")
	wordCount := flag.Int("words", 10, "Sets the word count for random mode")
	scores := flag.Bool("scores", false, "Display the highscores")
	filePath := flag.String("file", "", "The path to a custom file to load")
	flag.Parse()
	if *scores {
		db.ShowScores()
		return
	}

	var text string
	switch *mode {
	case "random":
		rg := randomtext.NewRandomGenerator(logger)
		text = rg.GenerateText(*wordCount)
	case "custom":
		if *filePath == "" {
			fmt.Println("Please enter a file path using --file")
			logger.Fatal("No path was specified")
		}
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println("unable to load file")
			logger.Fatal(err)
		}
		readValue, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("unable to read file contents")
			logger.Fatal(err)
		}
		text = string(readValue)
	default:
		text = "the quick brown fox jumps over the lazy dog"
	}

	execute(text, logger, db)
}

func execute(text string, logger *log.Logger, db *database.DB) {
	model := types.NewModel(text, logger, db)
	program := tea.NewProgram(model)
	program.Start()
}
