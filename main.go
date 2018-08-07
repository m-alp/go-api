package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Surname string `json:"last_name"`
	Age int `json:"age"`
	Password string `json:"-"`
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

type Response struct {
	Person
	TheName string `json:"the_name"`
}

func main()  {
	// create the middle handler
	middleware := func(handler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			startTime := time.Now()

			handler(writer, request)

			fmt.Printf("Exec Duration: %s", time.Since(startTime))

		}
	}

	http.HandleFunc("/", middleware(myHandler))

	fmt.Println("Listening to port 8080...")
	http.ListenAndServe(":8080", nil)
}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	bs, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	var p Person
	err := json.Unmarshal(bs, &p) // bytes array to struct
	// example error: send request with age as string (it's an int)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return // because there's return, the function is stopped
	}
	//fmt.Fprintf(writer, "Welcome %s %s \n", p.Name, p.Surname)

	var a map[string]interface{}
	json.Unmarshal(bs, &a)
	//fmt.Fprintf(writer, "Welcome %+v", a)

	response := Response{
		Person: p,
		TheName: p.FullName(),
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response) // write response in writer with json format
}


//func fetchGoogle(riter http.ResponseWriter, request *http.Request) {
//	response, _ := http.Get("https://www.google.com")
//
//	bs, _ := ioutil.ReadAll(request.Body)
//
//
//
//}