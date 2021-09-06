## Gowitness Command

Using {{target}}, {{rootDomain}} and {{domain}} ReconNess replace {{target}} for the target name, {{rootDomain}} for the rootdomain name and {{domain}} for the subdomain.

```
gowitness single -o /Content/screenshots/{{target}}/{{rootDomain}}/{{domain}}.png https://{{domain}}
```
## Gowitness Command for Docker

```
mkdir -p /app/Content/screenshots/{{target}}/{{rootDomain}}/ && /root/go/bin/gowitness single -o /app/Content/screenshots/{{target}}/{{rootDomain}}/{{domain}}.png https://{{domain}}
```

## Gowitness Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Gowitness/Script)

## Gowitness Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gowitness inside the docker

RUN wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
RUN GO111MODULE=on /usr/local/go/bin/go get -u github.com/sensepost/gowitness

# -------- End Agents dependencies -------- 
```
