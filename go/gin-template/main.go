package main

import (
	"log"
	"mtlogin/cmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	cmd.Execute()
}
