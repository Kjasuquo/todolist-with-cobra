package fortesting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

var Ab Data

// Data Manipulation: ----- My Data Structures ------

type Data struct {
	List   string `json:"list"`
	Status bool   `json:"status"`
}

var DataList []Data

func init() {
	Read()
}

// Read and Write: -------Reusable functions--------
func Read() {
	con, er := ioutil.ReadFile("task.csv")
	if er != nil {
		fmt.Println(er)
	}

	json.Unmarshal(con, &DataList)

}

func Write(s []Data) {
	json_data, erro := json.Marshal(s)
	if erro != nil {
		panic(erro)
	}

	e := ioutil.WriteFile("task.csv", json_data, 0666)
	if e != nil {
		fmt.Println(e)
	}
}

// My Functions: -----add, list, done, undone, cleanup -------

func (a *Data) AddFunc(s string) string {
	r := Data{List: s, Status: false}

	//Read()

	DataList = append(DataList, r)

	Write(DataList)

	return s + " added successfully"
}

func ListData() string {
	//Read()
	for i, v := range DataList {
		if v.Status == false {
			fmt.Println(i+1, v.List)
		}
	}

	return "list successfully printed"
}

func (a *Data) DoneTask(s string) string {
	y, _ := strconv.Atoi(s)
	//Read()
	for i, _ := range DataList {
		if y == (i + 1) {
			DataList[i].Status = true
		}
	}
	//fmt.Println(DataList)
	//fmt.Print("Successful\n")
	Write(DataList)

	return "Successful"
}

func (a *Data) UndoneItem(s string) string {
	y, _ := strconv.Atoi(s)
	//Read()
	for i, _ := range DataList {
		if y == (i+1) && DataList[i].Status == true {
			DataList[i].Status = false
		}
	}
	Write(DataList)
	return "Successful"
}

func Cleanup() string {
	var d []Data
	r := Data{}
	//Read()
	for _, value := range DataList {
		if value.Status == false {
			r.List = value.List
			r.Status = false

			d = append(d, r)
		}
	}

	Write(d)

	return "Successful"
}
