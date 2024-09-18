package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LeRoid-hub/Mensa-CLI/models"
)

func GetMenu(city string, mensa string) (models.Mensa, error) {
	http, err := http.Get("https://mensa.barfuss.email/mensa/" + city + "/" + mensa)
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
