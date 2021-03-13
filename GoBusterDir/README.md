## GoBusterDir Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

If we have the wordlist in ~/Desktop/tools/wordlist/directories.txt

```
gobuster dir -u https://{{domain}} -w ~/Desktop/tools/wordlist/directories.txt -z -k -r --wildcard 
```
## GoBusterDir Command for Docker

```
/root/go/bin/gobuster dir -u https://{{domain}} -w /app/content_discovery_all.txt -z -k -r --wildcard
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
RUN wget https://gist.githubusercontent.com/gorums/0a3a9d903e8e47fbff9d91097e19b4f8/raw/c81a34fe84731430741e0463eb6076129c20c4c0/content_discovery_all.txt

# -------- End Agents dependencies -------- 
```
