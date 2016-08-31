package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/jondot/runs"
)

func main() {
	files, _ := ioutil.ReadDir(".")
	things := []interface{}{}
	for _, file := range files {
		things = append(things, file)
	}

	grouped := runs.Detect(things, func(thing interface{}) int64 {
		return thing.(os.FileInfo).ModTime().Unix()
	}, func(a, b int64) bool {
		return math.Abs(float64(a-b)) < 60
	})
	for _, group := range grouped {
		fmt.Printf("-----\n")
		for _, file := range group {
			fi := file.(os.FileInfo)
			fmt.Printf("%v   %s\n", fi.ModTime(), fi.Name())
		}
	}

}
