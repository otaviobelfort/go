package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
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

type AddressEncoder struct {
	Cep         string
	Logradouro  string
	Complemento string
}

func main() {
	//cep := "01001000"
	//json_string := `{"cep":"123456","logradouro":"Rua das flores","complemento":"Casa 2"}`
	//fmt.Println(decoderString(json_string))
	//fmt.Println(encoderJson(AddressEncoder{Cep: "123456", Logradouro: "Rua das flores", Complemento: "Casa 2"}))
	Serve()
}

func BuscaCep(cep string) (Address, error) {
	response, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	//http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}

	// trazer a resposta como string
	//io.ReadAll(response.Body)

	if err != nil {
		return Address{}, err
	}
	defer response.Body.Close()

	var address Address
	if err := json.NewDecoder(response.Body).Decode(&address); err != nil {
		fmt.Printf("ERRO : %+v\n", err)
		return Address{}, err
	}

	//err = json.NewDecoder(response.Body).Decode(&address)
	//if err != nil {
	//	fmt.Printf("ERRO : %+v\n", err)
	//	return Address{}, err
	//}

	fmt.Printf("Address: %+v\n", address)
	return address, nil

}

func encoderJson(encoder AddressEncoder) string {
	encoderJson, err := json.Marshal(encoder)
	if err != nil {
		return ""
	}
	return string(encoderJson)

}

func decoderString(data_json string) AddressEncoder {
	var addressEncoder AddressEncoder
	//var addressEncode1 *AddressEncoder

	//addressEncode12 := AddressEncoder{Cep: "123456", Logradouro: "Rua das flores", Complemento: "Casa 2"}
	//addressEncode123 := &(AddressEncoder{Cep: "123456", Logradouro: "Rua das flores", Complemento: "Casa 2"})
	//*addressEncode1 = addressEncode12
	//println("VALOR : ", &addressEncode1)

	err := json.Unmarshal([]byte(data_json), &addressEncoder)
	//erro := json.NewEncoder(os.Stdout).Encode(data_json)
	if err != nil {
		return AddressEncoder{}
	}
	return addressEncoder
}

func Serve() {
	fmt.Println("Server ....")
	//http.HandleFunc("/cep", func(w http.ResponseWriter, r *http.Request) {
	//	address, _ := BuscaCep("65040430")
	//	fmt.Println("Address: ", address)
	//	fmt.Fprint(w, encoderJson(AddressEncoder{Cep: address.Cep, Logradouro: address.Logradouro, Complemento: address.Complemento}))
	//})

	http.HandleFunc("/cep1", GetCEP)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}
func GetCEP(w http.ResponseWriter, r *http.Request) {
	address, err := BuscaCep("01001000111")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(encoderJson(AddressEncoder{Cep: address.Cep, Logradouro: address.Logradouro, Complemento: address.Complemento})))
}
