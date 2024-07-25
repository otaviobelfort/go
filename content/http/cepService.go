package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Address1 struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type AddressEncoder1 struct {
	Cep         string
	Logradouro  string
	Complemento string
}

func main() {
	ServerFile()
}

func BuscaCep1(cep string) (*Address1, error) {
	http_time := http.Client{Timeout: time.Second}
	response, err := http_time.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var address Address1
	if err := json.NewDecoder(response.Body).Decode(&address); err != nil {
		fmt.Printf("ERRO : %+v\n", err)
		return nil, err
	}

	fmt.Printf("Address: %+v\n", address)
	return &address, nil

}
func BuscaCep3(cep string) (*Address1, error) {
	Newhttp := http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)

	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := Newhttp.Do(request)

	defer response.Body.Close()

	var address Address1
	if err := json.NewDecoder(response.Body).Decode(&address); err != nil {
		fmt.Printf("ERRO : %+v\n", err)
		return nil, err
	}

	fmt.Printf("Address: %+v\n", address)
	return &address, nil

}

// faça uma função que uma  requisicao http POST e Body pode ser um json de AddressEncoder e retorne um Address
func BuscaCepPost1(addressEncoder AddressEncoder1) (*Address1, error) {
	http_time := http.Client{Timeout: time.Second}
	// passe um Address1 no lugar no body

	response, err := http_time.Post("https://viacep.com.br/ws/01001000/json/",
		"application/json",
		bytes.NewBuffer([]byte(encoderJson1(addressEncoder))))

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected status 200, got %d", response.StatusCode)
	}

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var address Address1
	if err := json.NewDecoder(response.Body).Decode(&address); err != nil {
		fmt.Printf("ERRO : %+v\n", err)
		return nil, err
	}

	fmt.Printf("Address: %+v\n", address)
	return &address, nil

}

func encoderJson1(encoder AddressEncoder1) string {
	encoderJson, err := json.Marshal(encoder)
	if err != nil {
		return ""
	}
	return string(encoderJson)

}

func decoderString1(data_json string) AddressEncoder1 {
	var addressEncoder AddressEncoder1

	err := json.Unmarshal([]byte(data_json), &addressEncoder)
	if err != nil {
		return AddressEncoder1{}
	}
	return addressEncoder
}

func Server() {
	fmt.Println("Server ....")

	http.HandleFunc("/", GetCEP1)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}

func GetCEP1(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")
	address, err := BuscaCep1(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// verificar se o status da requisicao é 200
	if (w.WriteHeader(http.StatusOK)); err == nil {
		w.Write([]byte(encoderJson1(AddressEncoder1{Cep: address.Cep, Logradouro: address.Logradouro, Complemento: address.Complemento})))
	}
}

// SERVER MUX
func ServerFile() {
	fmt.Println("Server file....")
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

///-----------------------------------
//CONTEXT
