## Gau Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
echo '{{domain}}' | ./gau | grep "^https://{{domain}}"
```

## Gau Command for Docker

```
echo '{{domain}}' | /root/go/bin/gau | grep "^https://{{domain}}"
```

## Gau Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Gau/Script)

## Gau Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gau inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN /usr/local/go/bin/go get -u -v github.com/lc/gau

# -------- End Agents dependencies -------- 
```
