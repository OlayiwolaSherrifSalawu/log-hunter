package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Logstats struct {
	filename   string
	errorCount int
}

func (s *Logstats) CountError(lookup string) {
	file, err := os.Open(s.filename)

	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if strings.Contains(line, lookup) {
			s.errorCount += 1
		}
		if err != nil {
			break
		}

	}
	fmt.Printf("they are %d %s in file\n", s.errorCount, lookup)
}
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("invalid length of arguement ")
		return
	}

	lookup := args[1]

	s := Logstats{filename: "server.log"}
	s.CountError(lookup)
}
