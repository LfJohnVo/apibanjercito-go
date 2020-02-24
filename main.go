package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("*******COMENZANDO APLICACIÓN*********")
	/*fmt.Println("Seleccione una opción:")
	fmt.Println("[1] 1ZVBP8EM7D5276339")
	fmt.Println("[2] 3FADP4AJ7EM169194")
	fmt.Println("[3] Concurrente" )
	fmt.Println("[0] Salir")
	var n int
	num, err := fmt.Scanf("%v", &n)
	fmt.Println(err)
	fmt.Println("Tu numero es: ", num)*/

	
	consumo("3FADP4AJ7EM169194")


	/*jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}*/
	fmt.Println("Finalizando aplicación....")
}

func consumo(vin string){
	/* Comienza consumo de api*/
	response, err := http.Get("http://201.163.186.154:8080/BANJERCITO/wsvins/tests/v1/"+vin)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}