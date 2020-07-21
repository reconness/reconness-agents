package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Auth(url string, authApi string, username string, password string) string {

	values := map[string]string{"UserName": username, "Password": password}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", url+"/"+authApi, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func GetToken(jwt string) string {

	in := []byte(jwt)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return raw["auth_token"].(string)
}

func ExportSubdomains(url string, subdomainApi string, token string, filepath string) {

	req, err := http.NewRequest("GET", url+"/"+subdomainApi, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
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

func CsvtoTxt(existingCsv string, newTxt string) {
	var subDomains []string

	file, err := os.Open(existingCsv)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	for _, domain := range records {
		for _, subdomain := range domain {
			subDomains = append(subDomains, subdomain)
		}
	}

	newFile, err := os.OpenFile(newTxt, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(newFile)

	for _, data := range subDomains {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	newFile.Close()

}

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

//Subdomain to ip
func MassDns(subdomains string) {
	println("starting massdns")

	path := "/app/massdns"
	cmd := exec.Command(path+"/bin/massdns", "-r", path+"/lists/resolvers.txt", subdomains, "-o", "S", "-w", subdomains+".massdns")

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

	content := ReadFile(subdomains + ".massdns")
	for _, line := range content {
		fmt.Println(line)
	}

}

func main() {

	url := flag.String("b", "", "BaseUrl")

	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")

	authApi := flag.String("a", "", "Auth API")
	subdomainApi := flag.String("s", "", "Subdomain API")

	flag.Parse()

	randomFile := GenerateRandomString() + ".csv"

	// this is to ignore the cert
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Do the authentication and obtain the jwt
	jwt := Auth(*url, *authApi, *username, *password)
	// Get the token to allow us send auth request
	token := GetToken(jwt)
	// Export subdomains to a file
	ExportSubdomains(*url, *subdomainApi, token, "/tmp/"+randomFile)
	CsvtoTxt("/tmp/"+randomFile, "/tmp/"+randomFile+".txt")
	MassDns("/tmp/" + randomFile + ".txt")
}
