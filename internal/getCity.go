package internal

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

func GetState(state string) (string, error) {
	server := viper.Get("Server").(string)
	http, err := http.Get(server + "/state/" + state)
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
