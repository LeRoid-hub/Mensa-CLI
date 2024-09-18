package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetState(state string) (string, error) {
	fmt.Println("Get all called")

	http, err := http.Get("https://mensa.barfuss.email/state/" + state)
	if err != nil {
		return "", err
	}

	defer http.Body.Close()
	body, err := io.ReadAll(http.Body)
	if err != nil {
		return "", err
	}
	data := string(body)

	var cities []string
	err = json.Unmarshal([]byte(data), &cities)
	if err != nil {
		return "", err
	}

	var result string
	for _, city := range cities {
		result += city + ","

	}
	city := result[:len(result)-1]

	return city, nil
}
