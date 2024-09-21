package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeRoid-hub/Mensa-CLI/models"
	"github.com/spf13/viper"
)

func GetMenu(mensa string) (models.Mensa, error) {
	server := viper.Get("Server").(string)
	http, err := http.Get(server + "/mensa/" + mensa)
	if err != nil {
		return models.Mensa{}, err
	}

	defer http.Body.Close()

	body, err := io.ReadAll(http.Body)
	if err != nil {
		return models.Mensa{}, err
	}

	var data models.Mensa
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
		return models.Mensa{}, err
	}

	return data, nil
}

func GetSearchMenu(city string, mensa string) (models.Mensa, error) {
	return GetMenu(city + "/" + mensa)
}
