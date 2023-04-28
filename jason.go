package example_json

import (
	"encoding/json"
	"fmt"
)

type School struct {
	A int
	B int
	C int
}

type Person struct {
	School []School
	Name   string
	Age    int
}

func main() {
	jsonString := `{"school":[{"a":1,"b":2,"c":3},{"a":10,"b":20,"c":30}],"Name":"John Doe","Age":30}`
	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("JSON解析錯誤：", err)
		return
	}
	fmt.Println("姓名：", person.Name)
	fmt.Println("年齡：", person.Age)
	for i, s := range person.School {
		fmt.Printf("學校 %d:\n", i+1)
		fmt.Println("a：", s.A)
		fmt.Println("b：", s.B)
		fmt.Println("c：", s.C)
	}
}

