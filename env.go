package goenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) Load(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		c.strcet = strings.TrimSpace(s[:index])

		if len(c.strcet) == 0 {
			continue
		}

		value := strings.TrimSpace(s[index+1:])

		pos := strings.Index(value, "\t#")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " #")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, "\t//")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " //")
		if pos > -1 {
			value = value[0:pos]
		}

		if len(value) == 0 {
			continue
		}

		key := c.strcet
		c.Mymap[key] = strings.TrimSpace(value)
	}
}

func (c Config) Env(key string) string {
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}
