package pkg

import (
	"log"
)

func PrintIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
