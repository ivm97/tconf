package tconf

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

type Structure struct {
	config []string
}

type counter struct {
	value int64
}

func (c *counter) inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *counter) dec() {
	atomic.AddInt64(&c.value, -1)
}

func (c *counter) val() int64 {
	return atomic.LoadInt64(&c.value)
}

type insection map[string]any
type sections map[string]insection

// Read from path or io.Reader
func read(pathOrIo interface{}) (*[]byte, error) {
	if val, ok := pathOrIo.(string); ok {
		f, err := os.ReadFile(val)
		if err != nil {
			return nil, err
		}

		return &f, nil
	}

	if val, ok := pathOrIo.(*os.File); ok {
		f, err := os.ReadFile(val.Name())
		if err != nil {
			return nil, err
		}

		return &f, nil
	}

	return nil, fmt.Errorf("Can't read the file...")
}

func analyze(b *[]byte) (*sections, error) {
	if *b == nil {
		return nil, fmt.Errorf("nil data")
	}
	var (
		sect      counter
		dic       = make(sections)
		secdic    = make(insection)
		dchar     = []rune(string(*b))
		rbuff     []rune
		sectName  string = "default"
		key       string
		sectPoint int
		kvStart   int
		kvEnd     int
	)

	for i := 0; i < len(*b); i++ {

		switch dchar[i] {
		case '>':
			sectPoint = i + 1
			//secdic = make(map[string]any)
			secdic = make(insection)
			sect.inc()
		case '\n':
			if sect.val() == 1 {
				sect.dec()
				rbuff = dchar[sectPoint:i]
				sectName = strings.ReplaceAll(string(rbuff), " ", "")
				rbuff = rbuff[:0]

			}
			kvStart = i + 1

		case '~':
			rbuff = dchar[kvStart : i-1]
			key = strings.ReplaceAll(string(rbuff), " ", "")
			kvEnd = i + 1
			rbuff = rbuff[:0]

		case ';':

			rbuff = dchar[kvEnd:i]
			secdic[key] = valueType(string(rbuff))
			dic[sectName] = secdic
			rbuff = rbuff[:0]

		}

	}

	return &dic, nil

}

func valueType(target string) any {
	if r, err := strconv.Atoi(target); err != nil {
		return strings.ReplaceAll(target, " ", "")
	} else {
		return r
	}
}
