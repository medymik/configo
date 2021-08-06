package env

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Env struct {
	path string
}

func NewEnv(path string) *Env {
	return &Env{
		path: path,
	}
}

func getKeyValue(l string) (string, string) {
	key := ""
	value := ""
	is := true
	for _, ch := range l {
		if ch == '=' {
			is = false
		} else if is {
			key += string(ch)
		} else {
			value += string(ch)
		}
	}
	return key, value
}

func (e *Env) Load() {
	f, err := os.Open(e.path)
	if err != nil {
		log.Fatalf("Error %v", err.Error())
	}
	reader := bufio.NewReader(f)
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error %v", err.Error())
		}
		k, v := getKeyValue(string(l))
		os.Setenv(k, v)
	}
}
