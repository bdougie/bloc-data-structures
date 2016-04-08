package main

import (
	"fmt"
	"hash/fnv"
)

type HashClass struct {
	hash  uint32
	store [][2]string
}

func (h *HashClass) Insert(k string, v string) {
	if len(h.store) > 0 {
		for _, item := range h.store {
			if item[0] == k {
				item[1] = v
				fmt.Println(h.store)
				return
			}
		}
	}

	item := [2]string{k, v}
	h.store = append(h.store, item)
	fmt.Println(h.store)
}

func (h *HashClass) Get(k string) string {
	if h.store == nil {
		panic("Error: No values in the hash")
	}

	var v string

	for _, item := range h.store {
		if item[0] == k {
			v = item[1]
		}
	}

	return v
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
