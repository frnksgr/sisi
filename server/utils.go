package server

import (
	"log"
	"os"
	"strconv"
)

var logger = log.New(os.Stdout, "sisi", log.LstdFlags|log.Lshortfile)

func byteSize(s string) int64 {
	if len(s) > 0 {
		var factor int64
		switch f := s[len(s)-1]; f {
		case 'B':
			factor = 1
		case 'K':
			factor = 1024
		case 'M':
			factor = 1024 * 1024
		case 'G':
			factor = 1024 * 1024 * 1024
		}
		//n, err := strconv.Atoi(s[:len(s)-1])
		n, err := strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err == nil {
			return n * factor
		} else {
			logger.Println(err)
		}
	}
	return 0
}
