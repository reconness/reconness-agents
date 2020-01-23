## Ffuf Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target

If we have the wordlist in ~/Desktop/tools/wordlist/directories.txt

```
ffuf -w ~/Desktop/tools/wordlist/directories.txt -u https://{{domain}}/FUZZ
```

## Ffuf Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Ffuf/Script)

## Ffuf Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run ffuf inside the docker
RUN apt-get update && apt-get install -y wget && apt-get install -y git
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get github.com/ffuf/ffuf
RUN wget https://raw.githubusercontent.com/maurosoria/dirsearch/master/db/dicc.txt

# -------- End Agents dependencies -------- 
```

## Ffuf Command for Docker

```
cd /root/go/bin/ && ./ffuf -w /app/dicc.txt -u https://{{domain}}/FUZZ
```
