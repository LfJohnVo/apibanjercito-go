package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("*******COMENZANDO APLICACIÓN*********")

	vins := [2]string{"3FADP4AJ7EM169194", "1ZVBP8EM7D5276339"}
	tamaño := len(vins)

	for f := 0; f < tamaño; f++ {
	consumo(vins[f])
	}

	fmt.Println("Finalizando aplicación....")
}

func consumo(vin string) {
	/* Comienza consumo de api*/
	response, err := http.Get("http://201.163.186.154:8080/BANJERCITO/wsvins/tests/v1/" + vin)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		crearXml(data, vin)
	}
	//fmt.Println("Termina consumo")
}

func crearXml(data []byte, vin string) {
	//xmlString := string(data)
	fmt.Println("Resultado " +vin+ " : ", data)
	filename := vin+".xml"
	//file, _ := os.Create(filename)
	//xmlWriter := io.Writer(file)
	// Write `Hello, world!` to test.txt that can read/written by user and read by others
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}


}