## Naabu Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
./naabu -host {{domain}} -silent
```

## Naabu for Docker

```
cd /root/go/bin/ && ./naabu -host {{domain}} -silent
```

## Naabu Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Naabu/Script)

## Naabu Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run naabu inside the docker

RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get -v github.com/projectdiscovery/naabu/cmd/naabu

# -------- End Agents dependencies -------- 
```
