/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/kjasuquo/todolist/fortesting"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Long:  `Use this command to add task you want to carry out`,
	Run: func(cmd *cobra.Command, args []string) {
		//GetData(args)
		//ab := DataStructure{}
		//fmt.Println(GetData(args[0]))
		//addTodoList(args)
		fmt.Println(fortesting.Ab.AddFunc(args[0]))

	},
}

//var Data1 []byte

//
//func GetData(s string) string {
//	content, err := ioutil.ReadFile("task.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Printf("File contents: %s\n", content)
//
//	t := s + "\n"
//	Data1 = []byte(t)
//
//	content = append(content, Data1...)
//
//	er := ioutil.WriteFile("task.csv", content, 0666)
//	if er != nil {
//		fmt.Println(er)
//	}
//	//fmt.Println(Data1)
//
//	//In order to help with my done and undone commands
//	newContent, err := ioutil.ReadFile("task.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println(string(content))
//
//	var container []string
//	for _, record := range newContent {
//		container = append(container, string(record))
//	}
//	h := strings.Join(container, "")
//	j := strings.Split(h, "\n")
//
//	for _, v := range j {
//		M[v] = false
//	}
//
//	json_data, erro := json.Marshal(M)
//	if erro != nil {
//		//fmt.Println(err)
//	}
//
//	e := ioutil.WriteFile("done.csv", json_data, 0666)
//	if e != nil {
//		fmt.Println(e)
//	}
//
//	return "\n" + s + " successfully added"
//}

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
