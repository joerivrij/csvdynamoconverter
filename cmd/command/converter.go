package command

import (
	handler "csvdynamoconverter/pkg/impl"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/gookit/color.v1"
	"os/user"
)

func ConvertCsv() *cobra.Command {
	var (
		name      string
		filePath  string
		outFile   string
	)
	cmd := &cobra.Command{
		Use:   "convert",
		Short: "convert a csv",
		Long: `Allows you to convert a csv
- Name
- Filepath
- Outfile
`,
		Run: func(cmd *cobra.Command, args []string) {
			color.Green.Println("converting csv")
			if filePath == "" {
				color.Red.Println(fmt.Sprintf("filepath is empty"))
				return
			}
			if name == "" {
				color.Red.Println(fmt.Sprintf("name is empty"))
				return
			}
			if outFile == "" {
				usr, _ := user.Current()
				color.Yellow.Println(fmt.Sprintf("no outfile set, using to default %s/%s", usr.HomeDir, "Documents"))
				outFile = fmt.Sprintf("%s/%s", usr.HomeDir, "Documents")
			}

			outFile = fmt.Sprintf("%s/%s.json", outFile, name)
			handler.ConvertCsv(name, filePath, outFile)
		},
	}
	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the cluster")
	cmd.PersistentFlags().StringVarP(&filePath, "filepath", "f", "", "where to find the csv")
	cmd.PersistentFlags().StringVarP(&filePath, "outfile", "o", "", "where to store the json")

	return cmd
}