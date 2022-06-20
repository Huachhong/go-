package baselib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func TestHttp() {
	httpGet()
}

type Student map[string]interface{}

func httpGet() {
	//response, err := http.Get("https://www.baidu.com/?tn=sitehao123_15")
	response, err := http.Get("http://127.0.0.1:9090/userinfo")

	if err != nil {
		fmt.Println("handler err", err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var stu Student
	err = json.Unmarshal(body, &stu)
	if err != nil {
		fmt.Println("json format struct fail")
	}

	fmt.Printf("%#v", stu)
}

func httpPost() {

	url := "http://127.0.0.1:9090/user/edit"
	contentType := "application/json"
	json := `{"City":"广州"}`
	respon, err := http.Post(url, contentType, strings.NewReader(json))
	if err != nil {
		fmt.Println("http post fail")
	}
	defer respon.Body.Close()

}
