## Subfinder Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target

```
subfinder -d '{{domain}} -nW -silent'
```

## Subfinder Command for Docker

```
subfinder -d {{rootDomain}} -nW -silent
```

## Subfinder Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Subfinder/Script)

## Subfinder Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run subfinder inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN wget https://github.com/projectdiscovery/subfinder/releases/download/v2.4.5/subfinder_2.4.5_linux_amd64.tar.gz
RUN tar -xzvf subfinder_2.4.5_linux_amd64.tar.gz
RUN mv subfinder /usr/local/bin/

# -------- End Agents dependencies -------- 
```
