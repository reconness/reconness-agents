## GoBuster Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target.

If we have the wordlist in ~/Desktop/tools/wordlist/all.txt

```
go run main.go dns -d {{domain}} -w ~/Desktop/tools/wordlist/all.txt --wildcard
```
## GoBuster Command for Docker

```
cd /root/go/bin/ && ./gobuster dns -d {{domain}} -w /app/Subdomain.txt --wildcard
```

## GoBuster Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/GoBuster/Script)

## GoBuster Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gobuster inside the docker

RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
RUN . ~/.profile && go get github.com/OJ/gobuster
RUN wget https://raw.githubusercontent.com/gorums/WordLists/master/Subdomain.txt

# -------- End Agents dependencies -------- 
```
