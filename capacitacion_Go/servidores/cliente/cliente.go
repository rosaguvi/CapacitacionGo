package main

import (
	"net/http"	"net/url"
)


func main() {
	parametros := make(map[string]string)
	parametros["idProducto"] = "100"
	parametros["idUsusario"] = "3"
	nuevaURL := generarUrl("http", "localhost:5000", "/parametros", parametros)
	r, err := http.NewRequest ("GET", nuevaURL , nil)
	if err != nil {
		log.Println("Error en la Petición...")
		panic(err)
	}
	cliente := &http.Client{}
	response, errorResp := cliente.Do(r)
	if errorResp != nil {
		log.Println("Error en la respuesta...")
		panic(errorResp)
	}
	log.Println(response)
}

func generarUrl(protocolo , host, uri string , params map[string]string){
	nuevaURL , err := url.Parse(uri)
	if err != nil {
		log.Println("Error conviertiendo la URL")
		panic(err)
	}
	nuevaURL.Host = host
	nuevaURL.scheme = protocolo
	mapaParametros := nuevaURL.Query()
	for k , v := range params {
		mapaParametros.Add (k , v)
	}
	nuevaURL.RawQuery = mapaParametros.Encode()
	strUrl := nuevaURL.string()
	log.Println ("URL" , strUrl)
	return strUrl
}