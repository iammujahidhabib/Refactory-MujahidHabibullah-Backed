package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	NAME     string `json:"name"`
	PRICE    string `json:"price"`
	CATEGORY string `json:"category"`
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sort.Slice(records[1:], func(i, j int) bool {
		return records[1:][i][2] < records[1:][j][2]
	})

	return records
}
func createCSVtoJSON(filePath string) {
	csv_file, err := os.Open("01.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sort.Slice(records[1:], func(i, j int) bool {
		return records[1:][i][2] < records[1:][j][2]
	})

	var itm Item
	var items []Item

	length := len(records)
	fmt.Println(length)
	for i := 0; i < len(records); i++ {
		if i > 0 {
			itm.NAME = records[i][0]
			itm.CATEGORY = records[i][1]
			itm.PRICE = "Rp." + records[i][2]
			// fmt.Println(itm)
			items = append(items, itm)
		}
	}
	// fmt.Println(items)

	// Convert to JSON
	// j, _ := json.Marshal(items)
	// log.Println(string(j))

	// Convert to JSON
	json_data, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
		return
	}
	//print json data
	fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("01.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}
func main() {
	records := readCsvFile("01.csv")
	fmt.Println(records)
	createCSVtoJSON("01.csv")

}
