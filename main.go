package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/roniahmad/parking-app/bootstrap"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "the name of the file")
	flag.Parse()
}

func main() {
	app := bootstrap.NewApp()

	if filename == "" {
		fmt.Println("Please provide a filename using -f flag")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		if len(command) == 0 || command[0] == "" {
			fmt.Println("Empty or invalid command")
			continue
		}
		switch strings.ToLower(command[0]) {
		case "create_parking_lot":
			if len(command) < 2 {
				fmt.Println("Missing argument for create_parking_lot")
				continue
			}
			lotSize, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Printf("Invalid parking lot size: %v\n", err)
				continue
			}
			_ = app.LotController.CreateParkingLot(lotSize)
		case "park":
			if len(command) < 2 {
				fmt.Println("Missing argument for park")
				continue
			}

			number := strings.TrimSpace(command[1])
			if number == "" {
				fmt.Println("Missing argument for park")
				continue
			}
			_ = app.AllocController.Park(number)
		case "leave":
			if len(command) < 3 {
				fmt.Println("Missing argument for leave")
				continue
			}

			number := strings.TrimSpace(command[1])
			if number == "" {
				fmt.Println("Missing argument for leave")
				continue
			}
			hours, err := strconv.Atoi(command[2])
			if err != nil {
				fmt.Printf("Invalid parking hours: %v\n", err)
				continue
			}
			_ = app.AllocController.Leave(number, hours)

		case "status":
			_ = app.AllocController.Status()
		default:
			fmt.Println(scanner.Text(), "Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
