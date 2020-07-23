package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type columns struct {
	ID   string
	Text string
}

func main() {

	tableMap := make(map[string]string)

	// get column ID from StdIn
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Column Index to Fetch Data: ")
	column, _ := inputReader.ReadString('\n')
	columnID, err := strconv.Atoi(strings.TrimSuffix(column, "\n"))
	if err != nil {
		log.Fatal("Invalid column index")
	}

	// load csv file
	csvFile, _ := os.Open("text.csv")
	if err != nil {
		log.Fatal("Error occurred while loading csv file")
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	count := 0
	for {
		count++
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if count == 1 {
			continue
		}
		if len(line[0]) > 0 {
			tableMap[line[0]] = line[columnID]
		}
	}

	// write in JSON
	jsonString, _ := json.Marshal(tableMap)
	ioutil.WriteFile("text.json", jsonString, os.ModePerm)
}
