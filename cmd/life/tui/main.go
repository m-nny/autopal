package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/life"
)

func getBlinker() *life.GameState {
	return life.MustFromString(5, 5, `
		.....
		..#..
		..#..
		..#..
		.....
`)
}
func getBlock() *life.GameState {
	return life.MustFromString(4, 4, `
		....
		.##.
		.##.
		....
`)
}
func getBeehive() *life.GameState {
	return life.MustFromString(6, 5, `
		......
		..##..
		.#..#.
		..##..
		......
`)
}

func getToad() *life.GameState {
	return life.MustFromString(6, 6, `
		......
		...#..
		.#..#.
		.#..#.
		..#...
		......
`)
}

func getPulsar() *life.GameState {
	return life.MustFromString(17, 17, `
		.................
		.................
		....OOO...OOO....
		.................
		..O....O.O....O..
		..O....O.O....O..
		..O....O.O....O..
		....OOO...OOO....
		.................
		....OOO...OOO....
		..O....O.O....O..
		..O....O.O....O..
		..O....O.O....O..
		.................
		....OOO...OOO....
		.................
		.................
`)
}

func main() {
	logf, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	tickDuration := 100 * time.Millisecond
	p := tea.NewProgram(life.NewModel(getPulsar(), tickDuration))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
