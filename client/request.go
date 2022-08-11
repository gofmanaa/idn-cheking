package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

const Url = "https://post.refugeepass.nl/api/v1/appointment/get-alternative-options"

const ConfigFile = "config.toml"

type Conf struct {
	Cookie string `toml:"cookie"`
	Token  string `toml:"xsrf-token"`
}

type Payload struct {
	Day                Day   `json:"day"`
	Place              Place `json:"place"`
	AppointmentOptions struct {
		Appointment string `json:"appointment"`
	} `json:"appointment_options"`
	Info struct {
		TelephoneNumber string `json:"telephone_number"`
	} `json:"info"`
	Confirm struct {
	} `json:"confirm"`
}
type Day struct {
	Date   string `json:"date"`
	Amount string `json:"amount"`
}
type Place struct {
	Postcode string `json:"postcode"`
}

func NewPayload(date string) Payload {
	return Payload{
		Day:   Day{Date: date, Amount: "2"},
		Place: Place{Postcode: "9711"},
		AppointmentOptions: struct {
			Appointment string "json:\"appointment\""
		}{""},
		Info: struct {
			TelephoneNumber string "json:\"telephone_number\""
		}{""},
		Confirm: struct{}{},
	}
}

func GetConfig() (*Conf, error) {
	conf := &Conf{}
	d, err := toml.DecodeFile(ConfigFile, &conf)
	if err != nil {
		return nil, err
	}
	log.Println(d)

	return conf, nil
}

func PostRequest(log *logrus.Logger, date string, conf *Conf) {
	client := http.DefaultClient

	requestBody := createBody(date)

	req, _ := http.NewRequest("POST", Url, requestBody)

	createHeader(&req.Header, conf)

	log.Debugln("Header: %v", req.Header)
	log.Debugln("Body: %v", req.Body)

	resp, err := client.Do(req)

	if err != nil {
		log.Errorf("Request Failed: %s\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Update cookei and token in config file")
		log.Fatalf("Request Status: %s\n", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
	}

	ReadResponce(log, body)
}

func createHeader(header *http.Header, conf *Conf) {
	header.Add("authority", "post.refugeepass.nl")
	header.Add("accept", "*/*")
	header.Add("content-type", "application/json")
	header.Add("origin", "https://portaal.refugeepass.nl")
	header.Add("referer", "https://portaal.refugeepass.nl")
	header.Add("sec-ch-ua", `"Chromium";v="103", ".Not/A)Brand";v="99"`)
	header.Add("sec-ch-ua-mobile", `?0`)
	header.Add("sec-ch-ua-platform", `"Linux"`)
	header.Add("sec-fetch-dest", `empty`)
	header.Add("sec-fetch-mode", `cors`)
	header.Add("sec-fetch-site", `same-site`)
	header.Add("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36`)
	header.Add("x-requested-with", `XMLHttpRequest`)

	header.Add("cookie", conf.Cookie)
	header.Add("x-xsrf-token", conf.Token)

}

func createBody(date string) *bytes.Buffer {
	payload := NewPayload(date)
	postBody, _ := json.Marshal(payload)
	return bytes.NewBuffer(postBody)
}
