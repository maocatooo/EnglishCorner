package cmd

import (
	"EnglishCorner/server"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// document https://cobra.dev/

var rootCmd = &cobra.Command{
	Use:   "ec [run, init, import]",
	Short: "root command",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		server.InitData()
	},
}

var file string
var importCmd = &cobra.Command{
	Use:   "import [-f file]",
	Short: "import file",
	Run: func(cmd *cobra.Command, args []string) {

		if file == "" || !strings.HasSuffix(file, ".txt") {
			fmt.Println("error file for ", file)
		} else {
			server.Import(file)
		}
	},
}

func addCmd() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(importCmd)
}

func addVar() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "import words 'xxx.txt' Library name is 'xxx' ")

}

func init() {
	addCmd()
	addVar()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
