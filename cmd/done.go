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
		fmt.Println(Done(args[0]))
	},
}

var Notdone []string

var Workdone []byte

func Done(text string) string {
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

	donwork, err := ioutil.ReadFile("done.csv")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(j)-1; i++ {
		if ind == i+1 {
			workdone := j[i] + "\n"
			Workdone = []byte(workdone)
			donwork = append(donwork, Workdone...)
			er := ioutil.WriteFile("done.csv", donwork, 0666)
			if er != nil {
				fmt.Println(er)
			}
		}

		if ind != i+1 {
			Notdone = append(Notdone, j[i])
		}
	}

	nd, err := ioutil.ReadFile("done.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(nd))

	return "Successfully done tasks"
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
