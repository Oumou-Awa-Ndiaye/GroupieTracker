package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

	parts := strings.Split(center, "-")
    city := strings.Replace(parts[0], "_", "%20", -1)
    country := strings.Replace(parts[1], "_", "%20", -1)

func GetLocation(Artist) {
	url := "https://api.geoapify.com/v1/geocode/search?city=" + city + "&country=" + country + "&format=json&apiKey=118f1dbf888648258df3f09eb742819a"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
