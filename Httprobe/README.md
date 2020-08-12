## Httprobe Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo '{{domain}}' | httprobe
```

## Httprobe Command for Docker

```
echo '{{domain}}' | /root/go/bin/httprobe
```

## Httprobe Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Httprobe/Script)

## Httprobe Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run httprobe inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get github.com/tomnomnom/httprobe

# -------- End Agents dependencies --------  
```
