package mhw

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type dataScanner struct {
	lineScanner  *bufio.Scanner
	line         string
	fieldScanner *bufio.Scanner
	fieldIndex   int
}

func newDataScannerFromString(rawData string) *dataScanner {
	s := &dataScanner{}
	s.lineScanner = bufio.NewScanner(strings.NewReader(rawData))
	s.lineScanner.Split(bufio.ScanLines)
	s.fieldIndex = -1
	return s
}

func (s *dataScanner) scanRow() (ok bool) {
	s.fieldScanner = nil
	s.line = ""
	s.fieldIndex = -1
	for {
		if ok = s.lineScanner.Scan(); !ok {
			break
		}
		line := s.lineScanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "#"):
			continue
		}

		s.line = line
		s.fieldScanner = bufio.NewScanner(strings.NewReader(line))
		s.fieldScanner.Split(bufio.ScanWords)
		break
	}
	return ok
}

func (s *dataScanner) scanField() (ok bool) {
	if ok = s.fieldScanner.Scan(); ok {
		s.fieldIndex++
	}
	return ok
}

func (s *dataScanner) textValue() string {
	return s.fieldScanner.Text()
}

func (s *dataScanner) intValue() int {
	text := s.textValue()
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
	return value
}
