package main

type Logstats struct {
	filename   string
	errorCount int
}

func (s *Logstats) countError(filename string)
