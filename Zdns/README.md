## Zdns Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo "{{domain}},8.8.8.8" | go run zdns/main.go A
```

## Zdns Command for Docker
 
```
cd /app/zdns/zdns && echo '{{domain}},8.8.8.8' | /usr/local/go/bin/go run main.go A
```

## Zdns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Zdns/Script)

## Zdns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run zdns inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz 
RUN git clone https://github.com/zmap/zdns.git
RUN cd zdns && /usr/local/go/bin/go build

# -------- End Agents dependencies -------- 
```
