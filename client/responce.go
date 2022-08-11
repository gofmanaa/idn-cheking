package client

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Responce []struct {
	UniqueID     string `json:"unique_id"`
	Time         string `json:"time"`
	Date         string `json:"date"`
	Location     string `json:"location"`
	LocationData struct {
		UniqueID    string      `json:"unique_id"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Address     string      `json:"address"`
		Link        interface{} `json:"link"`
	} `json:"location_data"`
}

func ReadResponce(log *logrus.Logger, in []byte) {
	var out Responce
	err := json.Unmarshal(in, &out)
	if err != nil {
		log.Errorln(err)
	}
	log.Infof("Responce: %+v", out)
	if len(out) > 2 {
		log.Infoln("Go to https://portaal.refugeepass.nl/uk/make-an-appointment and fing appointment")
	} else {
		log.Infoln("Nothing found")
	}

}
