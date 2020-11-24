package impl

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/gookit/color.v1"
	"os"
	"strings"
)

func ConvertCsv(name, filePath, outputFile string) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//take the first element as the base and the rest as the json fields
	baseRecord := records[0]
	jsonRecords := records[1:]
	var typesOfRecords []string

	for j, base := range baseRecord {
		for i, char := range base {
			if string(char) == "(" {
				typesOfRecords = append(typesOfRecords, string(base[i+1]))
			}
		}
		splittedString := strings.Split(base, " ")[0]
		baseRecord[j] = splittedString
	}

	var csvToDynamoConvertedJson string
	csvToDynamoConvertedJson += fmt.Sprintf(`{"%s":[`, name)

	for j, jsonRecord := range jsonRecords {for i, record := range jsonRecord {
			if i == 0 {
				startLine := `{"PutRequest":{ "Item":{`
				csvToDynamoConvertedJson += startLine
			}
			if i != len(baseRecord)-1 {
				commaSeperated := fmt.Sprintf(`"%s":{ "%s": "%s"},`, baseRecord[i], typesOfRecords[i], record)
				csvToDynamoConvertedJson += commaSeperated
			} else {
				lastElement := fmt.Sprintf(`"%s":{ "%s": "%s"}`, baseRecord[i], typesOfRecords[i], record)
				csvToDynamoConvertedJson += lastElement
			}

			if i == len(baseRecord)-1 && j != len(jsonRecords)-1{
				endLineWithComma := `}}},`
				csvToDynamoConvertedJson += endLineWithComma
			} else if i == len(baseRecord)-1 && j == len(jsonRecords)-1 {
				endLine := `}}}]}`
				csvToDynamoConvertedJson += endLine
			}
		}
	}

	openedFile, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer openedFile.Close()

	outputFromWrite, err := openedFile.WriteString(csvToDynamoConvertedJson)
	if err != nil {
		return
	}

	color.Green.Println(fmt.Sprintf("finished writing %d bytes", outputFromWrite))
	color.Green.Println(fmt.Sprintf("file written to %s", outputFile))

	openedFile.Sync()

	color.Green.Println("use the line below to deploy to AWS dynamo")
	fmt.Println(fmt.Sprintf("aws dynamodb batch-write-item --request-items file://%s", outputFile))

	return
}
