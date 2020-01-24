## ffuf Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

If we have the wordlist in ~/Desktop/tools/wordlist/directories.txt

```
ffuf -w ~/Desktop/tools/wordlist/directories.txt -u https://{{domain}}/FUZZ
```

## ffuf Command for Docker

```
cd /root/go/bin/ && ./ffuf -w /app/dicc.txt -u https://{{domain}}/FUZZ
```

## ffuf Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Ffuf/Script)

## ffuf Dockerfile Entry

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
RUN wget https://gist.githubusercontent.com/gorums/0a3a9d903e8e47fbff9d91097e19b4f8/raw/c81a34fe84731430741e0463eb6076129c20c4c0/content_discovery_all.txt

# -------- End Agents dependencies -------- 
```


