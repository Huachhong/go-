package baselib

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayWeb(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./baselib/template.html")
	if err != nil {
		fmt.Println("create Template failf", err)
		return
	}
	tmpl.Execute(w, "golang")
}

func TestTemplate() {
	http.HandleFunc("/", sayWeb)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("http server failf, err:", err)
		return
	}

}
