package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//ReadFile reads the content of a file and returns an array of it's content
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
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return content
}

//OneForAll executes the OneForAll python script
func OneForAll(domain string) {
	cmd := exec.Command("python3", "/app/OneForAll/oneforall.py", "--target", domain, "--path", "tmp", "run")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmd.Wait()

}

//ParseData reads the output of the OneForAll script and prints it to stdout
func ParseData() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			content := ReadFile(cwd + "/tmp/" + file.Name())
			for _, line := range content {
				println(line)
			}
		}
	}
	os.RemoveAll("tmp/")
}

func main() {
	domain := flag.String("d", "", "Domain")
	flag.Parse()

	OneForAll(*domain)
	ParseData()

}
