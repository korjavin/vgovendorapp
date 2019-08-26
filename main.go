package main

import (
	"github.com/korjavin/vgovendorlib"
	log "github.com/sirupsen/logrus"
)

func main() {
	b, _ := vgovendorlib.Something()
	log.Printf("vim-go %s", b)
}
