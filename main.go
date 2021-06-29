package main

import (
	"fmt"
	lider "liderscraping/web"
)

func main() {
	fmt.Println("Iniciando")
	liderWeb := lider.LiderScraping{}
	liderWeb.Init()
}
