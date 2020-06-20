package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
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
func OneForAll(domain string, outputDir string) {
	cmd := exec.Command("python3", "/app/OneForAll/oneforall.py", "--target", domain, "--path", outputDir, "run")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmd.Wait()

}

//ParseData reads the output of the OneForAll script and prints it to stdout
func ParseData(outputDir string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(outputDir + "/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".txt") {
			content := ReadFile(cwd + "/" + outputDir + "/" + file.Name())
			for _, line := range content {
				println(line)
			}
		}
	}
	os.RemoveAll(outputDir + "/")
}

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijklmnopqrstuvwxyz" +
			"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"

	return str
}

func main() {
	domain := flag.String("d", "", "Domain")
	flag.Parse()

	outputDir := GenerateRandomString()

	OneForAll(*domain, outputDir)
	ParseData(outputDir)

}
