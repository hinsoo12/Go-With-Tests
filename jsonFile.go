package main

import (
	"fmt"
	"encoding/json"
	"log"
)


type GOIntership struct{
	Company string
	DateJoined string `json:"joined"`
	NameOfEmployers []string
	Position string
	Payment bool `json:"payment, notpaid"`
	NumberOfEmployers int `json:"total"`
}

func jsonMarshaling(result string) string{
		
var intership = []GOIntership{
	{Company: "2F-Capital",DateJoined:"07/03/2022",NameOfEmployers:[]string{"Henok Mengistu","Nahom Asnake","Gadisa Teka","Hawi Girma"},
     Position: "Intership",Payment: false, NumberOfEmployers: 7},
}

/*
data, err :=json.Marshal(intership)  
if err !=nil{
	log.Fatalf("JSON Conversion is failed : %s", err)
}
fmt.Printf("%s\n", data)
*/


list , errors :=json.MarshalIndent(intership,""," ")
if errors !=nil{
	log.Fatalf("Marshaling is failed : %s", errors)
}
fmt.Println("\nAll required information of 2F-Capital Go Intership Employers")
fmt.Printf(" %s\n", list)

return result

}