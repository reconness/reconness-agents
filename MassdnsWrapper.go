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
func ExportSubdomains(url string, subdomainApi string, token string, extraArguments string) {

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

	MassDns(tmpFile.Name(), extraArguments)
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
func MassDns(subdomains string, extraArguments string) {

	//path := "/app/massdns"
	path := "/home/bossman/tools/massdns"
	optArguments := strings.Fields(extraArguments)
	arguments := []string{path + "/bin/massdns", "-r", path + "/lists/resolvers.txt", "-o", "S", subdomains, "-w", subdomains + ".massdns"}

	if len(optArguments) > 0 {
		for _, optArg := range optArguments {
			arguments = append(arguments, optArg)
		}
	}

	cmd := exec.Command(arguments[0], arguments[1:]...)

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

	subdomainApi := flag.String("s", "", "Subdomain API")
	optArgs := flag.String("o", "", "Optional arguments used in Massdns -r, -o, -s and -w are already used. Make sure to put it in qoutes e.g \"--type A -c 3\"")

	flag.Parse()

	if *username == "" {
		fmt.Println("Please provide a username")
		os.Exit(1)
	} else if *password == "" {
		fmt.Println("Please provide a password")
		os.Exit(1)
	} else if *subdomainApi == "" {
		fmt.Println("Please provide a subdomain API endpoint")
		os.Exit(1)
	}

	authApi := "api/Auth/Login"

	// this is to ignore the cert
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Do the authentication and obtain the jwt
	jwt := Auth(*url, authApi, *username, *password)
	// Get the token to allow us send auth request
	token := GetToken(jwt)
	// Export subdomains to a file & run massdns
	ExportSubdomains(*url, *subdomainApi, token, *optArgs)
}
