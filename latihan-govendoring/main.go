package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"io/ioutil"
	"log"
	"net/http"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main(){

	for{
		water := h8HelperRand.RandomInt(1, 100)
		wind := h8HelperRand.RandomInt(1, 100)

		waterStatus := "aman"
		if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else if water > 8 {
			waterStatus = "bahaya"
		}

		windStatus := "aman"
		if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else if wind > 15 {
			windStatus = "bahaya"
		}

		



		data := map[string]interface{}{
			"water": water,
			"wind": wind,
		}
	
		reqestJson,err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}
	
		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqestJson))
	
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
	
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}	
		defer res.Body.Close()
	
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		
		// print post result
		fmt.Println("Post Result:")
		fmt.Println(string(body))

		// print water and wind status
		fmt.Println("Water Status:", waterStatus)
		fmt.Println("Wind Status:", windStatus)

		// wait for 15 seconds before the next post
		time.Sleep(15 * time.Second)
	
	}

}