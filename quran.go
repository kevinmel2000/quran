package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Ayah   `json:"data"`
}

type Ayah struct {
	Number        int    `json:"number"`
	Text          string `json:"text"`
	NumberInSurah int    `json:"numberInSurah"`
	Juz           int    `json:"juz"`
	Surah         Surah  `json:"surah"`
}

type Surah struct {
	Number                 int    `json:"number"`
	Name                   string `json:"name"`
	EnglishName            string `json:"englishName"`
	EnglishNameTranslation string `json:"englishNameTranslation"`
	NumberOfAyahs          int    `json:"numberOfAyahs"`
	RevelationType         string `json:"revelationType"`
}

func printAyah(surah, ayah string) {
	url := fmt.Sprintf("http://api.alquran.cloud/ayah/%s:%s", surah, ayah)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf(
		"Surah: %s (%s) - %s\n",
		response.Data.Surah.EnglishName,
		response.Data.Surah.Name,
		response.Data.Surah.EnglishNameTranslation,
	)
	fmt.Printf(
		"Ayah #%d in Juz %d\n",
		response.Data.NumberInSurah,
		response.Data.Juz,
	)
	fmt.Println(response.Data.Text)
}
