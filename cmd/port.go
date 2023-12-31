/*
Copyright © 2023 Bestony bestony@linux.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer fc.Close()

	content, err := io.ReadAll(fc)
	if err != nil {
		panic(err)
	}

	return content
}

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port [zip file path]",
	Short: "port flomo zip file to target dest",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		zf, err := zip.OpenReader(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer zf.Close()

		for _, file := range zf.File {
			if strings.Contains(file.Name, "index") {
				fmt.Printf("=%s\n", file.Name)
				fmt.Printf("%s\n\n", readAll(file)) // file content
			}

		}

		fmt.Println("port called")
	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
