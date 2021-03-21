package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Databases struct {
	Database []Database `json:"database_classification"`
}

type Database struct {
	DbName           string `json:"database_name"`
	DbOwner          string `json:"database_owner"`
	DbClassification string `json:"classification"`
}

/*
type Employees struct {
	Employee []Employee `json:"employee_list`
}
*/

type Employee struct {
	row_id       int
	user_id      string
	user_state   string
	user_manager string
}

func main() {
	fmt.Println("Starting parse files...")

	jsonFile, err := os.Open("database.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Opened database.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var databases Databases

	json.Unmarshal(byteValue, &databases)

	for i := 0; i < len(databases.Database); i++ {
		fmt.Println("Database name: " + databases.Database[i].DbName)
		fmt.Println("Database owner: " + databases.Database[i].DbOwner)
		fmt.Println("Database classification: " + databases.Database[i].DbClassification)
	}

	lines, err := ReadCsv("user_manager.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(lines)
	/*
		for _; line := range lines {
			data := CsvLine (
				row_id: 	line[0],
				user_id:	 	line[1],
				user_state:		line[2],
				user_manager:	line[3],
			)
			fmt.Println(data.DbLineId + " " + data.DbUserId)
		}
	*/
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}
