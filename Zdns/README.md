## Zdns Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo "{{domain}},8.8.8.8" | go run zdns/main.go A
```

## Zdns Command for Docker
 
```
cd /app/zdns/zdns && echo "{{domain}},8.8.8.8" | /usr/local/go/bin/go run main.go A
```

## Zdns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Zdns/Script)

## Zdns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run zdnsinside the docker
RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN git clone https://github.com/zmap/zdns.git
RUN . ~/.profile && cd zdns && go build

# -------- End Agents dependencies -------- 
```
