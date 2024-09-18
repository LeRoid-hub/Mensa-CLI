package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMensen(city string) (string, error) {
	fmt.Println("Get all called")

	http, err := http.Get("https://mensa.barfuss.email/city/" + city)
	if err != nil {
		return "", err
	}

	defer http.Body.Close()
	body, err := io.ReadAll(http.Body)
	if err != nil {
		return "", err
	}
	data := string(body)

	var mensen []string
	err = json.Unmarshal([]byte(data), &mensen)
	if err != nil {
		return "", err
	}

	var result string
	for _, mensa := range mensen {
		result += mensa + ","

	}
	mensa := result[:len(result)-1]

	return mensa, nil
}
