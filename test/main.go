package main

import (
	"fmt"
	"github.com/RestartFU/Go-TheHive"
	"log"
)

func main() {
	r := hive.NewPlayerRequest(hive.TimeAll(), hive.GameSkyWars(), "xMyma")
	resp, err := r.Do()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
