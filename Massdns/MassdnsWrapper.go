package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//Auth authenticates to reonness
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

//GetToken Retrieves the JWT for authentication
func GetToken(jwt string) string {

	in := []byte(jwt)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return raw["auth_token"].(string)
}

//ExportSubdomains exports subdomains to a temporary file
func ExportSubdomains(url string, subdomainApi string, token string) {

	req, err := http.NewRequest("GET", url+"/"+subdomainApi, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	subDomains := strings.Split(string(bodyBytes), ",") //csv format to txt

	tmpFile, err := ioutil.TempFile(os.TempDir(), "massdns-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	for _, data := range subDomains {
		text := []byte(data + "\n")
		if _, err = tmpFile.Write(text); err != nil {
			log.Fatal("Failed to write to temporary file", err)
		}
	}

	MassDns(tmpFile.Name())
}

//ReadFile reads the file and saves to an array
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

//MassDns Execute MassDNS
func MassDns(subdomains string) {

	path := "/app/massdns"
	cmd := exec.Command(path+"/bin/massdns", "-r", path+"/lists/resolvers.txt", subdomains, "-o", "S", "-w", subdomains+".massdns")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
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

	// this is to ignore the cert
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Do the authentication and obtain the jwt
	jwt := Auth(*url, *authApi, *username, *password)
	// Get the token to allow us send auth request
	token := GetToken(jwt)
	// Export subdomains to a file & run massdns
	ExportSubdomains(*url, *subdomainApi, token)
}
