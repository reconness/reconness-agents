## Shuffledns Command

Using {{rootDomain}} ReconNess replace {{rootDomain}} for the root domain.
To allow run this tool we need to have a resolvers.txt and massdns installed

```
shuffledns -d {{rootDomain}} -r ~/resolvers.txt -w /app/all.txt -massdns ~/massdns/bin/massdns -silent"
```

## Shuffledns Command for Docker

```
cd /root/go/bin/ && ./shuffledns -d {{rootDomain}} -r /app/resolvers.txt -w /app/all.txt -massdns /app/massdns/bin/massdns -silent
```

## Shuffledns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Shuffledns/Script)

## Shuffledns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run shuffledns inside the docker

RUN apt-get update && apt-get install -y git wget
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN cd /app && wget https://raw.githubusercontent.com/reconness/reconness-agents/master/resolvers.txt
RUN . ~/.profile && go get -u -v github.com/projectdiscovery/shuffledns/cmd/shuffledns

# -------- End Agents dependencies -------- 
```
