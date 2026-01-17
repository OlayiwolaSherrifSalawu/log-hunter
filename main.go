package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Logstats struct {
	filename   string
	errorCount int
}

func (s *Logstats) CountError(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		} else if strings.Contains(line, "ERROR") {
			s.errorCount += 1
		}
	}
	defer file.Close()
	fmt.Println("they are " + strconv.Itoa(s.errorCount) + " in file")
}
func main() {
	s := Logstats{}
	s.CountError("server.log")
}
