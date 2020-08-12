## Dnsprobe Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo {{domain}} | ./dnsprobe -r A
```

## Dnsprobe Command for Docker

```
echo {{domain}} | /root/go/bin/dnsprobe -r A
```

## Dnsprobe Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Dnsprobe/Script)

## Dnsprobe Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run dnsprobe inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get -u -v github.com/projectdiscovery/dnsprobe

# -------- End Agents dependencies -------- 
```
