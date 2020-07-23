package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// user input file path
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	filePath, _ := inputReader.ReadString('\n')

	jsonFile, err := os.Open(strings.TrimSuffix(filePath, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		log.Fatal("Error unmarshalling json")
	}

	s := strings.Split(filePath, "/")
	fileName := s[len(s)-1]

	// write in JSON
	jsonString, _ := json.Marshal(objmap)
	ioutil.WriteFile(fileName+".json", jsonString, os.ModePerm)
}
