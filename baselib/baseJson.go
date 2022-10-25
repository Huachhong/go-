package baselib

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age,string"`
	Niubility bool   `json:"Niubility"`
}

func TestJson() {
	b := []byte(`{"name":"小虎","age":18,"niubility":true}`)
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}
	}
	fmt.Printf("%#v", i)
}
