/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
	"strings"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long:  `After a task, use this to mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		Done(args[0])
	},
}

func Done(text string) {
	var Dbyte []byte
	ind, _ := strconv.Atoi(text)
	var s []string
	content, err := ioutil.ReadFile("task.csv")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(content))

	for _, record := range content {
		s = append(s, string(record))
	}
	h := strings.Join(s, "")
	j := strings.Split(h, "\n")

	for i := 0; i < len(j)-1; i++ {
		if ind != i+1 {
			fmt.Println(i+1, j[i])
			t := j[i] + "\n"
			Dbyte = []byte(t)

		}
	}
	//fmt.Println(j)

	content = append(content, Dbyte...)

	er := ioutil.WriteFile("db.csv", content, 0666)
	if er != nil {
		fmt.Println(er)
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
