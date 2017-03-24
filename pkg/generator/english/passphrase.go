package english

import (
	"math/rand"
	"os"
	"time"

	"bufio"

	"strings"

	"fmt"

	"log"

	"github.com/karlockhart/phrase-gen/pkg/generator/methods"
	"github.com/karlockhart/phrase-gen/pkg/util"
	"github.com/spf13/viper"
)

type passPhrase struct {
	graph *methods.Graph
}

func readDataFile(filename string) (words []string, n int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, strings.TrimSpace(s.Text()))
		n++
	}
	return
}
func (p *passPhrase) getRandomVerb() string {
	return p.getRandomWord("verbs")
}

func (p *passPhrase) getRandomWord(k string) string {

	if (util.Cache.Get(k)) == nil {
		dir := viper.GetString("dictionaries.directory")
		words, count, err := readDataFile(fmt.Sprintf("%s/%s.list", dir, k))
		if err != nil {
			log.Fatal(err)
		}
		util.Cache.Put(k, &words, count)
		log.Println("Disk")
		return words[rand.Intn(count)]
	}
	ce := util.Cache.Get(k)
	words := *ce.Words
	count := ce.Count
	log.Println("Cache")
	return words[rand.Intn(count)]

}

func NewPassPhrase() *passPhrase {
	rand.Seed(time.Now().UTC().UnixNano())
	e := new(passPhrase)
	e.graph = methods.NewGenerator()

	log.Println(e.getRandomVerb())
	log.Println(e.getRandomVerb())

	//r := e.graph.GetRoot()
	//r.AddNode()
	return e
}
