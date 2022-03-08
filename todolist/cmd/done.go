/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
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
	Long: `After a task, use this to mark a task as done
This temporarily deletes the tasks you mark as done, only displaying undone tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Done(args[0]))
	},
}

var M = make(map[string]bool)

func Done(text string) string {
	//var test string
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

	con, er := ioutil.ReadFile("done.csv")
	if er != nil {
		fmt.Println(er)
	}
	json.Unmarshal(con, &M)
	//fmt.Println(M)

	for index, value := range j {
		for item, _ := range M {
			if ind == index+1 && value == item {
				M[value] = true
			}
			if item == "" {
				M[item] = true
			}
		}
	}

	//fmt.Println(M)

	json_data, err := json.Marshal(M)
	if err != nil {
		//fmt.Println(err)
	}

	e := ioutil.WriteFile("done.csv", json_data, 0666)
	if er != nil {
		fmt.Println(e)
	}

	for item, status := range M {
		if status == false {
			fmt.Println(item)
		}
	}

	return "\n" + text + " has been successfully done"
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
