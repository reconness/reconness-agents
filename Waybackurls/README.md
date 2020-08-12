## Waybackurls Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo '{{domain}}' | ./waybackurls | grep "^https://{{domain}}"
```

## Waybackurls Command for Docker

```
echo '{{domain}}' | /root/go/bin/waybackurls | grep "^https://{{domain}}"
```

## Waybackurls Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Waybackurls/Script)

## Waybackurls Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run waybackurls inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get github.com/tomnomnom/waybackurls

# -------- End Agents dependencies -------- 
```
