## Dnsprobe Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo {{domain}} | ./dnsprobe -r A
```

## Dnsprobe Command for Docker

```
cd /root/go/bin/ && echo {{domain}} | ./dnsprobe -r A
```

## Dnsprobe Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Dnsprobe/Script)

## Dnsprobe Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run dnsprobe inside the docker

RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get -u -v github.com/projectdiscovery/dnsprobe

# -------- End Agents dependencies -------- 
```
