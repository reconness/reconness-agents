package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	var content []string

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return content
}

func OneForAll(domain string) {
	println("starting one for all")

	cmd := exec.Command("python3", "/root/app/OneForAll/oneforall.py", "--target", domain, "--path", "tmp", "run")

	println(cmd.String())
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	cmd.Wait()

}

func CopyOneForAll(dir string) {

	files, err := ioutil.ReadDir(dir + "tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			content := ReadFile(dir + "tmp/" + file.Name())
			for _, line := range content {
				println(line)
			}
		}
		os.RemoveAll(dir + "tmp/")
	}
}

func main() {
	domain := flag.String("d", "", "Domain")
	flag.Parse()

	OneForAll(*domain)
	CopyOneForAll("/root/app/OneForAll/")

}
