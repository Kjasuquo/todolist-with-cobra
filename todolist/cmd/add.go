/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Long:  `Use this command to add task you want to carry out`,
	Run: func(cmd *cobra.Command, args []string) {

		addToList([][]string{{"Task"}, {"Algo"}, {"Agile"}})
		//fmt.Println("add called")
	},
}

func addToList(s [][]string) {
	task := s

	csvfile, err := os.Create("task.csv")
	if err != nil {
		log.Fatalln(err)
	}
	cswriter := csv.NewWriter(csvfile)
	for _, t := range task {
		_ = cswriter.Write(t)
	}
	cswriter.Flush()
	csvfile.Close()
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
