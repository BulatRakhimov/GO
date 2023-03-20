package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	Calculate("Россия, г. Москва, Cлавянский бульвар д.1", "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43", 62000, 43, 35, 180, 1)
}

func Calculate(addrFrom string, addrTo string, ves int, dlina int, shirina int, visota int, tip int) ([]byte, error) {
	// токен авторизации
	// addrFrom string, addrTo string, ves int, dlina int, shirina int, visota int, tip int  ///([]itog, err error)
	//addrFrom = "Россия, г. Москва, Cлавянский бульвар д.1"
	//addrTo = "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"
	//ves = 62000
	//dlina = 43
	//shirina = 35
	//visota = 180
	//tip = 1

	dannie := "grant_type=client_credentials&client_id=EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI&client_secret=PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	otvet1, err := http.Post("https://api.edu.cdek.ru/v2/oauth/token", "application/x-www-form-urlencoded", bytes.NewBufferString(dannie))
	if err != nil {
		fmt.Println(err)
	}
	defer otvet1.Body.Close()

	var otvet_auth map[string]interface{}
	json.NewDecoder(otvet1.Body).Decode(&otvet_auth)

	token_tipa, ok := otvet_auth["access_token"].(string)
	if !ok {
		fmt.Println("Oshibka, invalid token")
	}
	//fmt.Println(token_tipa)

	zapros := map[string]interface{}{
		"from_location": map[string]interface{}{
			"address": addrFrom,
		},
		"to_location": map[string]interface{}{
			"address": addrTo,
		},
		"packages": []map[string]interface{}{
			{
				"weight": ves,
				"length": dlina,
				"width":  shirina,
				"height": visota,
				"type":   tip,
			},
		},
	}
	zapros_json, _ := json.Marshal(zapros)

	zapros2, err := http.NewRequest("POST", "https://api.edu.cdek.ru/v2/calculator/tarifflist", bytes.NewBuffer(zapros_json))
	if err != nil {
		fmt.Println(err)
	}

	zapros2.Header.Set("Content-Type", "application/json")
	zapros2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token_tipa))

	klient := &http.Client{}
	otvet2, oshibka := klient.Do(zapros2)
	if oshibka != nil {
		fmt.Println(oshibka)
	}
	defer otvet2.Body.Close()

	var itog map[string]interface{}
	json.NewDecoder(otvet2.Body).Decode(&itog)

	itog_json, _ := json.MarshalIndent(itog, "", "  ")
	fmt.Println(string(itog_json))
	return itog_json, oshibka
}
