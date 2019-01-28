package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileStat struct {
	path string
	os.FileInfo
}

func main() {
	dstSep := flag.String("s", "-", "path separator in target file")
	includeDir := flag.Bool("d", true, "include dir path in flattened file name?")
	flag.Parse()

	as := flag.Args()
	if len(as) != 2 {
		fmt.Println("Error: Missing input dir and output dir. They should be the last two arguments.")
		flag.Usage()
		return
	}

	if isExists(as[0]) != DIR {
		fmt.Println("Error: source dir does not exist: ", as[0])
		return
	}

	if isExists(as[1]) != DIR {
		fmt.Println("dest dir does not exist: ", as[1])
		//fmt.Println("Creating dir:", as[1])
		err := os.MkdirAll(as[1], 0766)
		if err != nil {
			fmt.Println("\nError creating dir ", as[1], ":", err)
			return
		}
		fmt.Println("Created:", as[1])
	}

	//var files []string
	var files []FileStat
	err := filepath.Walk(as[0], func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, FileStat{path, info})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		//fmt.Println(file.path)
		copyFile(file, as[0], as[1], *dstSep, *includeDir)
	}
}

func copyFile(srcFs FileStat, src, dst string, dstSep string, includeDir bool) {
	pathSep := string(filepath.Separator)
	if strings.HasSuffix(src, pathSep) {
		src = src[:len(src)-1]
	}
	if strings.HasSuffix(dst, pathSep) {
		dst = dst[:len(dst)-1]
	}

	sourceFile := srcFs.path
	var dstFile string
	if includeDir {
		dstFile = strings.Replace(sourceFile, src+pathSep, "", 1)
		dstFile = strings.Replace(dstFile, pathSep, dstSep, -1)
		dstFile = dst + pathSep + dstFile
	} else {
		dstFile = dst + pathSep + srcFs.Name()
	}

	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(dstFile, input, srcFs.Mode())
	if err != nil {
		fmt.Println("Error creating", dstFile)
		fmt.Println(err)
		return
	}

}

const (
	NOPE = iota
	FILE
	DIR
)

func isExists(path string) int {
	if stat, err := os.Stat(path); err == nil {
		if stat.IsDir() {
			return DIR
		}
		return FILE
	}
	return NOPE
}
