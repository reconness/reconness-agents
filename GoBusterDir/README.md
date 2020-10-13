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

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get github.com/OJ/gobuster
RUN wget https://gist.githubusercontent.com/gorums/0a3a9d903e8e47fbff9d91097e19b4f8/raw/c81a34fe84731430741e0463eb6076129c20c4c0/content_discovery_all.txt

# -------- End Agents dependencies -------- 
```
