/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as undone",
	Long: `A task that was not done but mistakenly marked as done can be marked undone back with this command
It is used to retrieve a task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Undone(args[0]))
	},
}

func Undone(text string) string {
	//var test string
	ind, _ := strconv.Atoi(text)
	var s []string
	content, err := ioutil.ReadFile("task.csv")
	if err != nil {
		fmt.Println(err)
	}

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

	for index, value := range j {
		for item, _ := range M {
			if ind == index+1 && value == item {
				M[value] = false
			}
			if item == "" {
				M[item] = true
			}
		}
	}

	json_data, err := json.Marshal(M)
	if err != nil {
		fmt.Println(err)
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

	return "\n" + "Task has been undeleted"
}
func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
