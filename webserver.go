package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
		)

type Person struct {
	Uid int
	First_name string
	Last_name string
	Age  int
	Email string
	Reference string
	ServerResponse bool
}


func index(w http.ResponseWriter, r *http.Request) {





	var person Person

	body, _ := ioutil.ReadAll(r.Body)
	
	xml.Unmarshal(body, &person)
	person.ServerResponse = true

	responseXML, _ := xml.Marshal(person) 
	
	
	fmt.Fprintf(w, string(responseXML))
	///log.Println(w,string(responseXML))


	log.Println(string("Marshaled xml: \n"))
	log.Println(person.Uid)
	log.Println(person.First_name)
	log.Println(person.Last_name)
	log.Println(person.Age)
	log.Println(person.Email)
	log.Println(person.Reference)


}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}
