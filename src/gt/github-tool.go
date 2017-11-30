package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Class struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (class Class) String() string {
	return fmt.Sprintf("Id=%d Name=%s", class.Id, class.Name)
}

func main() {
	// Pleaer change your data !
	g_user := "CHANGE-ME"
	p_token := "CHANGE-ME"

	url := "https://api.github.com/users/" + g_user + "/repos?per_page=10000"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "token "+p_token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))

	var classes []Class
	err = json.Unmarshal(body, &classes)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(classes); i++ {
		fmt.Println(classes[i])
	}
}
