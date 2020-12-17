package sms77api

import (
	"log"
	"reflect"
	"strconv"
)

func getMapKeys(mymap map[interface{}]interface{}) []interface{} {
	//mymap := make(map[int]string)
	keys := make([]interface{}, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}

	return keys
}

func pickMapByKey(needle interface{}, haystack interface{}) (interface{}, interface{}) {
	mapIter := reflect.ValueOf(haystack).MapRange()

	for mapIter.Next() {
		log.Printf("%#v", needle)
		log.Printf("%#v", mapIter.Key())

		if needle == mapIter.Key() {
			return needle, mapIter.Value()
		}
	}

	return nil, nil
}

func inArray(needle interface{}, haystack interface{}) bool {
	slice := reflect.ValueOf(haystack)
	c := slice.Len()

	for i := 0; i < c; i++ {
		if needle == slice.Index(i).Interface() {
			return true
		}
	}

	return false
}

func toUint(id string, bitSize int) uint64 {
	n, err := strconv.ParseUint(id, 10, bitSize)

	if nil == err {
		return n
	}

	return 0
}
