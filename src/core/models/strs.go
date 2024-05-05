package models

import (
	"database/sql/driver"
	"strings"
)

type Strs []string

func (s *Strs) Scan(val interface{}) error {
	*s = strings.Split(val.(string), "|")
	return nil
}

func (s Strs) Value() (driver.Value, error) {
	str := strings.Join(s, "|")
	return str, nil
}

func (s Strs) Contains(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (s Strs) Add(str string) Strs {
	if s.Contains(str) {
		return s
	}
	return append(s, str)
}
