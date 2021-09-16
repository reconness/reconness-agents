## ffuf Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

If we have the wordlist in ~/Desktop/tools/wordlist/directories.txt

```
ffuf -w ~/Desktop/tools/wordlist/directories.txt -u https://{{domain}}/FUZZ
```

## ffuf Command for Docker

```
/root/go/bin/ffuf -w /app/Content/wordlists/dir_enum/default.txt -u https://{{domain}}/FUZZ
```

## ffuf Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Ffuf/Script)

## ffuf Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run ffuf inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get github.com/ffuf/ffuf

# -------- End Agents dependencies -------- 
```
