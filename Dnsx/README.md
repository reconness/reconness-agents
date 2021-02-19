## Dnsx Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo {{domain}} | ./dnsx -silent -a -resp
```

## Dnsx Command for Docker

```
echo {{domain}} | dnsx -silent -a -resp
```

## Dnsx Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Dnsx/Script)

## Dnsx Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run dnsx inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN wget https://github.com/projectdiscovery/dnsx/releases/download/v1.0.1/dnsx_1.0.1_linux_amd64.tar.gz
RUN tar -xzvf dnsx_1.0.1_linux_amd64.tar.gz
RUN mv dnsx /usr/local/bin/

# -------- End Agents dependencies -------- 
```
