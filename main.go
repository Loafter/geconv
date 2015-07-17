package main

import (
	"log"
	"os/exec"
	"fmt"
)

func main() {
	fmt.Println("This autorun writen by Andrew Rose GE HC Russia")
	cmd := exec.Command("java", "-jar","DICOMVWR.CAB")
	err := cmd.Start()
	if err != nil {
		down:=exec.Command("explorer","https://java.com/ru/download/")
		down.Start()
		log.Fatal(err)
	}
}
