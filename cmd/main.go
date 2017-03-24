package main

import (
	_ "github.com/karlockhart/phrase-gen/pkg/config"
	"github.com/karlockhart/phrase-gen/pkg/generator/english"
)

func main() {
	english.NewPassPhrase()

}
