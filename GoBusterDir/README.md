## GoBusterDir Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

If we have the wordlist in ~/Desktop/tools/wordlist/directories.txt

```
gobuster dir -u https://{{domain}} -w ~/Desktop/tools/wordlist/directories.txt -s 200,204 -z -k -r --wildcard 
```
## GoBusterDir Command for Docker

```
/root/go/bin/gobuster dir -u https://{{domain}} -w /app/Content/wordlists/dir_enum/default.txt -s 200,204 -z -k -r --wildcard
```

## GoBusterDir Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/GoBusterDir/Script)

## GoBusterDir Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gobuster inside the docker

RUN wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
RUN /usr/local/go/bin/go install github.com/OJ/gobuster/v3@latest

# -------- End Agents dependencies -------- 
```
