package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "io"
    "bytes"
//     "reflect"
//     "strings"
    "net/url"
)

const (
    API_KEY = "6de53479dfc406f2cca00b7d0d5f9e83"
)

type InputArtist struct {
    input string `json:"input_char"`
}

type ArtistDetail struct {
    Name string `json:"name"`
    URL string `json:url`
}


// func GetJson(url string, target interface{}) error {
//     // will update func for get data
//
// }

type ArtistSearch struct {
    Name  string `json:"Name"`
    Url string `json:"Url"`
}

func main() {
     log.Println("Hello There..!")
     http.HandleFunc("/search", func(rw http.ResponseWriter, req* http.Request){
        // Input from postman
          reqBody, _ := ioutil.ReadAll(req.Body)
          keyVal := make(map[string]string)
          json.Unmarshal(reqBody, &keyVal) // check for errors
        input_search := keyVal["input_char"]

        // Url define
        url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.search&artist=%s&api_key=%s&format=json",url.QueryEscape(input_search),API_KEY)
           fmt.Println(url)
        r, er := http.Get(url)
        if er != nil {
            log.Fatal(er)
        }

        // Get Response Formatted
        responseData, err := ioutil.ReadAll(r.Body)
        if err != nil{
            log.Fatal(err)
        }

        reqs := struct {
		Body io.Reader
        }{bytes.NewReader(responseData)}
        fmt.Println(reqs)
        var settings interface{}
        if err := json.NewDecoder(reqs.Body).Decode(&settings); err != nil {
            panic(err)
        }


        list := settings.(map[string]interface{})
        results_data := list["results"].(map[string]interface{})
        artist_matches := results_data["artistmatches"].(map[string]interface{})
        artist_data := artist_matches["artist"]
        slice := artist_data.([]interface{})
        dynamic_value := slice[0].(map[string]interface{})
//             fmt.Println(reflect.TypeOf(dynamic_value["url"]))

        jsonString := ArtistSearch{Name: dynamic_value["name"].(string), Url: dynamic_value["url"].(string)}
        byteArray, err := json.MarshalIndent(jsonString, "", " ")
        if err != nil {
            fmt.Println(err)
        }
        rw.Write([]byte(string(byteArray)))
     })
    http.ListenAndServe(":6000" , nil)
}