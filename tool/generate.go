// +build ignore

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var name string
	flag.StringVar(&name, "o", "", "output")
	flag.Parse()

	var err error
	var output *os.File = os.Stdout
	if name != "" {
		output, err = os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
	}

	resp, err := http.Get("http://www.unicode.org/Public/vertical/revision-17/VerticalOrientation-17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	var from, to rune
	var prop string

	fmt.Fprintln(output, `package utr50

var table = []struct {
	from rune
	to   rune
	prop prop
}{`)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		token := strings.SplitN(line, ";", 2)
		if len(token) != 2 {
			continue
		}
		runes := strings.SplitN(token[0], "..", 2)
		u, err := strconv.ParseUint(strings.TrimSpace(runes[0]), 16, 64)
		if err != nil {
			continue
		}
		from = rune(u)
		if len(runes) > 1 {
			u, err = strconv.ParseUint(strings.TrimSpace(runes[1]), 16, 64)
			if err != nil {
				continue
			}
			to = rune(u)
		} else {
			to = from
		}
		prop = strings.TrimSpace(token[1])
		fmt.Fprintf(output, "\t{0x%X, 0x%X, %s},\n", from, to, prop)
	}
	fmt.Fprintln(output, "}")
}
