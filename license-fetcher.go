package main

import (
	"bufio"
	"fmt"
	"io"
	"license-fetcher/utils"
	"log"
	"os"
)

func main() {
	f, e := os.Open("./resources/repos.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, _, e := reader.ReadLine()
		if e != nil {
			if e != io.EOF {
				log.Println("ReadLine fail", e)
			}
			break
		}
		pkg := string(line)
		license, repo, e := utils.FetchForGo(pkg)
		if e != nil {
			fmt.Println(pkg)
		} else {
			fmt.Printf("%s,%s,%s\n", pkg, license, repo)
		}
	}

}
