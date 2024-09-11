package tconf

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	data sections
}

// Create new tconf file!
func New() *Structure {
	return &Structure{}
}

// Open tconf file!
func Open(pathOrIo interface{}) (*Config, error) {
	b, err := read(pathOrIo)
	if err != nil {
		return nil, err
	}
	tc, err := analyze(b)
	if err != nil {
		return nil, err
	}
	return &Config{data: *tc}, nil
}

func (c *Config) From(section string) *insection {
	if v, ok := c.data[section]; ok {
		return &v
	}

	return nil
}

func (in *insection) Get(key string) (interface{}, bool) {
	if v, ok := (*in)[key]; ok {
		return v, ok
	}

	return nil, false
}

// Add section to new config
func (s *Structure) AddSection(name string) {
	s.config = append(s.config, ">"+name)
}

// Add key and value to new config
func (s *Structure) AddKeyValue(key string, value any) {
	s.config = append(s.config, fmt.Sprintf("   %s ~ %v;", key, value))
}

// Save new config in file!
func (s *Structure) Save(filePath string) error {
	res := strings.Join(s.config, "\n")
	return os.WriteFile(filePath, []byte(res), 0666)
}
