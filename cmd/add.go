/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Long:  `Use this command to add task you want to carry out`,
	Run: func(cmd *cobra.Command, args []string) {
		//GetData(args)
		//ab := DataStructure{}
		GetData(args[0])
		//addTodoList(args)

	},
}

var Data []byte

var Fer []string

//type Items struct {
//	List DataStructure
//}
//
//type DataStructure struct {
//	item   []string
//	status bool
//}

//func GetData(s string) {
//	//r.item = s
//	//r.status = true
//	json_data, err := json.Marshal(s)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(json_data)
//
//	content, err := ioutil.ReadFile("task.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(content)
//
//	content = append(content, json_data...)
//
//	er := ioutil.WriteFile("task.csv", content, 0666)
//	if er != nil {
//		fmt.Println(er)
//	}
//	fmt.Println(content)
//
//}

func GetData(s []string) {
	//file := DataStructure{
	//	item:   s,
	//	status: false,
	//}
	content, err := ioutil.ReadFile("task.csv")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("File contents: %s\n", content)

	for i := 0; i < len(s); i++ {
		t := s[i] + "\n"
		Data = []byte(t)
	}

	content = append(content, Data...)

	er := ioutil.WriteFile("task.csv", content, 0666)
	if er != nil {
		fmt.Println(er)
	}
	//fmt.Println(Data)

	for _, v := range content {
		Fer = append(Fer, string(v))
	}
	b := strings.Join(Fer, "")
	fmt.Println(b)

	for _, v := range b {
		fmt.Println(v)

	}

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