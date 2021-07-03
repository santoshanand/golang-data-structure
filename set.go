//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

package main

import "fmt"

type Item string

type ItemSet struct {
	items map[string]bool
}

func (s *ItemSet) Add(t string) ItemSet {
	if s.items == nil {
		s.items = make(map[string]bool)
	}
	_, hasItem := s.items[t]

	if !hasItem {
		s.items[t] = true
	}
	return *s
}

func (s *ItemSet) Clear() {
	(*s).items = make(map[string]bool)
}

func main() {
	item := ItemSet{}

	item.Add("Hello")
	item.Add("Hello1")
	item.Add("Hello1")

	fmt.Println(item)
}
