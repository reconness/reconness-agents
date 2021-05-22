## Nuclei Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

If we have the templates in /root/nuclei-templates

```
nuclei -u {{domain}} -t /root/nuclei-templates -silent
```

## Nuclei Command for Docker
 
```
/root/go/bin/nuclei -u {{domain}} -t /root/nuclei-templates -silent
```

## Nuclei Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Nuclei/Script)

## Nuclei Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run nuclei inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz 
RUN GO111MODULE=on /usr/local/go/bin/go get -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei 
RUN /root/go/bin/nuclei -update-templates

# -------- End Agents dependencies -------- 
```
