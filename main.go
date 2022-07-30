package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func isAppended(f *os.File) (bool, error) {
	if _, err := f.Seek(-2, io.SeekEnd); err != nil {
		fmt.Println("seek error:", err)
		return false, err
	}

	suffix, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("read all error:", err)
		return false, err
	}
	if bytes.HasPrefix(suffix, []byte("#")) {
		return true, nil
	}
	return false, nil
}

func handlerDir(p string, fn func(p string, entry os.FileInfo)) {
	dir, err := os.ReadDir(p)
	if err != nil {
		log.Fatal("read dir error:", err)
	}
	rand.Seed(time.Now().UnixNano())
	for _, entry := range dir {
		fInfo, err := entry.Info()
		if err != nil {
			log.Fatal("get file info error:", err)
		}
		fn(filepath.Join(p, entry.Name()), fInfo)
	}
	fmt.Println("It's all done.")
}

func appendN(p string, entry os.FileInfo) {
	if strings.HasPrefix(entry.Name(), ".") {
		fmt.Println("skip file:", entry.Name())
		return
	}
	f, err := os.OpenFile(p, os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open file error:", err)
		fmt.Println("skip file:", entry.Name())
		return
	}
	stat, err := f.Stat()
	if err != nil {
		fmt.Println("get file info error:", err)
		fmt.Println("skip file:", entry.Name())
		return
	}
	if stat.Size() <= 0 {
		fmt.Println("skip too small file:", entry.Name())
		return
	}
	appended, err := isAppended(f)
	if err != nil {
		fmt.Println("skip file:", entry.Name())
		return
	}
	if appended {
		fmt.Println(entry.Name() + " no need to append")
		return
	}
	if _, err = f.WriteString(fmt.Sprintf("#%d", rand.Intn(10))); err != nil {
		fmt.Println("write string error:", err)
		fmt.Println("skip file:", entry.Name())
	} else {
		fmt.Println(entry.Name() + " append success")
	}
}

func recoverFile(p string, fInfo os.FileInfo) {
	if strings.HasPrefix(fInfo.Name(), ".") {
		fmt.Println("skip file:", fInfo.Name())
		return
	}
	f, err := os.OpenFile(p, os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("open file error:", err)
		fmt.Println("skip file:", fInfo.Name())
		return
	}
	appended, err := isAppended(f)
	if err != nil {
		fmt.Println("skip file:", fInfo.Name())
		return
	}
	if !appended {
		fmt.Println(fInfo.Name() + " no need for recovery")
		return
	}

	if err = f.Truncate(fInfo.Size() - 2); err != nil {
		fmt.Println("truncate error:", err)
		fmt.Println("skip file:", fInfo.Name())
	} else {
		fmt.Println(fInfo.Name() + " recover success")
	}
}

func usage() {
	fmt.Printf("Usage: %s [-h|--help] [-m|--modify] [-r|--recover] [--modify-One] [--recover-One] \n", os.Args[0])
}

func main() {
	args := os.Args
	switch args[1] {
	case "-h", "--help":
		usage()
	case "-m", "--modify":
		handlerDir(args[2], appendN)
	case "-r", "--recover":
		handlerDir(args[2], recoverFile)
	case "--modify-One":
		appendN(args[2], getFInfo(args[2]))
	case "--recover-One":
		recoverFile(args[2], getFInfo(args[2]))
	default:
		log.Fatal("args error")
	}
}

func getFInfo(path string) os.FileInfo {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal("open file error:", err)
	}
	fInfo, err := f.Stat()
	if err != nil {
		log.Fatal("open file error:", err)
	}
	return fInfo
}
