package baselib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func TestHttp() {
	httpGet()
}

func httpGet() {
	response, err := http.Get("https://www.baidu.com/?tn=sitehao123_15")
	if err != nil {
		fmt.Println("handler err", err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%v", string(body))
}
