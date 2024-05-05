package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CatFactResponse struct {
    Fact string `json:"fact"`
}

const URL = "https://catfact.ninja/fact"

func GetRandomCatFact() string {
    //Send GET request to api
    response, err := http.Get(URL)
    if err != nil {
        fmt.Println("ERROR: ", err)
        return ""
    }
    defer response.Body.Close()

    //Decode JSON response
    var catFact CatFactResponse
    err = json.NewDecoder(response.Body).Decode(&catFact)
    if err != nil {
        fmt.Println("Error decoding JSON: ", err)
        return ""
    }

    return catFact.Fact
}
