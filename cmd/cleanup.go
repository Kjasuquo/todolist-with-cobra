/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/kjasuquo/todolist/fortesting"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up done tasks",
	Long: `Clean up your task after it has been done.
This hard deletes your done tasks and you cannot get it back `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fortesting.Cleanup())
	},
}

//func Cleanup() string {
//	var holdIt []string
//	con, er := ioutil.ReadFile("done.csv")
//	if er != nil {
//		fmt.Println(er)
//	}
//	json.Unmarshal(con, &M)
//	//fmt.Println(M)
//
//	for i, v := range M {
//		if v == false {
//			holdIt = append(holdIt, i)
//		}
//	}
//	//fmt.Println(holdIt)
//
//	for _, v := range holdIt {
//		t := v + "\n"
//		r := []byte(t)
//		Data1 = append(Data1, r...)
//	}
//
//	err := ioutil.WriteFile("task.csv", Data1, 0666)
//	if er != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println(Data1)
//
//	return "All done tasks have been deleted"
//}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
