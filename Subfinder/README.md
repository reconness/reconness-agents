## Subfinder Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target

```
subfinder -d '{{domain}}'
```

## Subfinder Command for Docker

```
cd /root/go/bin/ && ./subfinder -d {{domain}}
```

## Subfinder Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Subfinder/Script)

## Subfinder Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run subfinder inside the docker
RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get -u github.com/projectdiscovery/subfinder/cmd/subfinder

# -------- End Agents dependencies -------- 
```
