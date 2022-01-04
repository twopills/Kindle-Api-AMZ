package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func formatterMonth(m time.Month) string {

	var month int = int(m)
	ITMonths := []string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "dicembre"}
	for i, value := range ITMonths {
		if i == month {
			return value
		}
	}
	return ""
}

func TakeDate() string {
	y, _, d := time.Now().Date()
	day := strconv.Itoa(d)
	var m time.Month
	month := formatterMonth(m)
	year := strconv.Itoa(y)
	return day + " " + month + " " + year
}

func GetAllBooks() Structure {

	url := Url
	method := "POST"

	payload := strings.NewReader(Payload_GetElem)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("authority", SiteNoHttps)
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("dnt", "1")
	req.Header.Add("user-agent", User_Agent)
	req.Header.Add("client", "MYX")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("origin", Site)
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", Api)
	req.Header.Add("accept-language", "en-US,en;q=0.9,it;q=0.8")
	req.Header.Add("cookie", Cookie)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return parseValueToObj(body)
}

func DeleteBooks(asin string) {
	url := Url
	method := "POST"

	payload := strings.NewReader(Payload1 + asin + Payload2)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", SiteNoHttps)
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("dnt", "1")
	req.Header.Add("user-agent", User_Agent)
	req.Header.Add("client", "MYX")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("origin", Site)
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", Api)
	req.Header.Add("accept-language", "en-US,en;q=0.9,it;q=0.8")
	req.Header.Add("cookie", Cookie)

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

func parseValueToObj(b []byte) Structure {
	var msg Structure

	err := json.Unmarshal(b, &msg)
	if err != nil {
		log.Fatalln(err)
	}
	return msg
}

func changeToCollection() {
	// todo
}

func selectElementFromDate(date string) {
	books := GetAllBooks().OwnerShipData.Items
	for _, b := range books {
		if b.AcquiredDate != date {
			DeleteBooks(b.Asin)
		}
		if b.AcquiredDate == date {
			// todo
		}
	}
}
