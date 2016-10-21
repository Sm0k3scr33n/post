package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Uid int
	First_name string
	Last_name string
	Age  int
	Email string
	Reference string
}

func main() {

	person := &Person{1, "Ben", "Gabbard", 30, "user@domain.com", "Weave"}
	buf, _ := xml.Marshal(person)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8080", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))

}
