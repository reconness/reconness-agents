## Dnsx Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo {{domain}} | ./dnsx -silent -a -resp
```

## Dnsx Command for Docker

```
echo '{{domain}}' | /root/go/bin/dnsx -silent -a -resp
```

## Dnsx Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Dnsx/Script)

## Dnsx Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run dnsx inside the docker

RUN wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
RUN GO111MODULE=on /usr/local/go/bin/go get -v github.com/projectdiscovery/dnsx/cmd/dnsx

# -------- End Agents dependencies -------- 
```
