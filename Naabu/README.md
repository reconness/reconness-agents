## Naabu Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
./naabu -host {{domain}} -silent
```

## Naabu for Docker

```
/root/go/bin/naabu -host {{domain}} -silent
```

## Naabu Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Naabu/Script)

## Naabu Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run naabu inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get -v github.com/projectdiscovery/naabu/v2/cmd/naabu

# -------- End Agents dependencies -------- 
```
