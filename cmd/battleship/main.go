package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/c0dect/battleship"
)

var (
	defaultInputFile  = "input.txt"
	defaultOutputFile = "output.txt"

	errParse = errors.New("error in parsing the input")
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flagset := flag.NewFlagSet("battleship", flag.ExitOnError)
	var (
		inputFile  = flagset.String("input", defaultInputFile, "input file for running the battleship simulator.")
		outputFile = flagset.String("output", defaultOutputFile, "output file for the battleship simulator.")
	)

	if err := flagset.Parse(args); err != nil {
		return err
	}

	// define the input file.
	inpFile, err := os.Open(*inputFile)
	if err != nil {
		return errors.New("could not open the input file")
	}
	defer inpFile.Close()

	// define the output file.
	var outFile *os.File
	outFile, err = os.Open(*outputFile)
	if err != nil {
		outFile, err = os.Create(*outputFile)
		if err != nil {
			return errors.New("could not open/create the output file")
		}
	}
	defer outFile.Close()

	// create the game object.
	game, err := parseInputToCreateGame(inpFile)
	if err != nil {
		return err
	}

	// run the game.
	game.SimulateGame()

	game.PrintResult(outFile)

	return nil
}

func parseInputToCreateGame(r io.Reader) (*battleship.Game, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	mString := scanner.Text()
	m, err := strconv.Atoi(mString)
	if err != nil {
		log.Print(err)
		return nil, errParse
	}
	scanner.Scan()

	var pos string
	scanner.Scan()
	pos = scanner.Text()
	p1c, err := parseShipPositions(pos)
	if err != nil {
		return nil, errParse
	}

	scanner.Scan()
	pos = scanner.Text()
	p2c, err := parseShipPositions(pos)
	if err != nil {
		return nil, errParse
	}
	scanner.Scan()

	scanner.Scan()
	pos = scanner.Text()
	t1c, err := parseMisPositions(pos)
	if err != nil {
		return nil, errParse
	}
	scanner.Scan()
	pos = scanner.Text()
	t2c, err := parseMisPositions(pos)
	if err != nil {
		return nil, errParse
	}

	p1 := battleship.CreatePlayer("Player1", m, p1c, t1c)
	p2 := battleship.CreatePlayer("Player2", m, p2c, t2c)

	return &battleship.Game{
		Player1: p1,
		Player2: p2,
	}, nil
}

func parseShipPositions(posString string) ([]battleship.Coordinates, error) {
	coords := make([]battleship.Coordinates, 0)
	strTokens := strings.Split(posString, ",")
	for _, v := range strTokens {
		xyTokens := strings.Split(v, ":")
		c1, err := strconv.Atoi(xyTokens[0])
		if err != nil {
			return nil, errParse
		}
		c2, err := strconv.Atoi(xyTokens[1])
		if err != nil {
			return nil, errParse
		}
		coords = append(coords, battleship.Coordinates{c1, c2})
	}
	return coords, nil
}

func parseMisPositions(posString string) ([]battleship.Coordinates, error) {
	coords := make([]battleship.Coordinates, 0)
	strTokens := strings.Split(posString, ":")
	for _, v := range strTokens {
		xyTokens := strings.Split(v, ",")
		c1, err := strconv.Atoi(xyTokens[0])
		if err != nil {
			return nil, errParse
		}
		c2, err := strconv.Atoi(xyTokens[1])
		if err != nil {
			return nil, errParse
		}
		coords = append(coords, battleship.Coordinates{c1, c2})
	}
	return coords, nil
}
