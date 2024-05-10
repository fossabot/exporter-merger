package main

import (
	"github.com/ne-bknn/exporter-merger/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := internal.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
