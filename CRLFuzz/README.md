## CRLFuzz Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.


```
crlfuzz -u https://{{domain}} -s
```

## CRLFuzz Command for Docker

```
/root/go/bin/crlfuzz -u https://{{domain}} -s
```

## CRLFuzz Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/CRLFuzz/Script)

## CRLFuzz Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run CRLFuzz inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get -v github.com/dwisiswant0/crlfuzz/cmd/crlfuzz

# -------- End Agents dependencies -------- 
```
