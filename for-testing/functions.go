package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetData(s string) string {
	var Data1 []byte
	content, err := ioutil.ReadFile("task.csv")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("File contents: %s\n", content)

	for i := 0; i < len(s); i++ {
		t := s + "\n"
		Data1 = []byte(t)
	}

	content = append(content, Data1...)

	er := ioutil.WriteFile("task.csv", content, 0666)
	if er != nil {
		fmt.Println(er)
	}
	//fmt.Println(Data1)

	return s + " successfully added"
}

func Read() string {
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
		fmt.Println(i+1, j[i])
	}
	return "list successfully printed"
}
