package main

import (
	commandCtl "csvdynamoconverter/cmd/command"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dynamoconvertercli",
		Short: "Converts csv to dynamojson",
		Long: `Converts a csv to dynamodb useable json. The command for uploading is printed in the end'`,
	}
)

func main() {
	rootCmd.AddCommand(
		commandCtl.ConvertCsv(),
	)
	err := rootCmd.Execute()
	if err != nil {
		color.Red.Println(err)
	}
}