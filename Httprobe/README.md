## Httprobe Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo '{{domain}}' | httprobe
```

## Httprobe Command for Docker

```
cd /root/go/bin/ && echo '{{domain}}' | ./httprobe
```

## Httprobe Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Httprobe/Script)

## Httprobe Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run httprobe inside the docker
RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get github.com/tomnomnom/httprobe

# -------- End Agents dependencies --------  
```
