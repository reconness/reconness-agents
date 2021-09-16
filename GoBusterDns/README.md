## GoBusterDns Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target.

If we have the wordlist in ~/Desktop/tools/wordlist/all.txt

```
gobuster dns -d {{domain}} -w ~/Desktop/tools/wordlist/all.txt --wildcard -z
```

## GoBusterDns Command for Docker

```
/root/go/bin/gobuster dns -d {{domain}} -w /app/Content/wordlists/subdomain_enum/default.txt --wildcard -z
```

## GoBusterDns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/GoBusterDns/Script)

## GoBusterDns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gobuster inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get github.com/OJ/gobuster

# -------- End Agents dependencies -------- 
```
