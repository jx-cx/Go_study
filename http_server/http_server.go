package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(file)
}
func f2(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("OK"))

}
func main() {
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
