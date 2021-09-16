## Shuffledns Command

Using {{rootDomain}} ReconNess replace {{rootDomain}} for the root domain.
To allow run this tool we need to have a resolvers.txt and massdns installed

```
shuffledns -d {{rootDomain}} -r ~/resolvers.txt -w /app/all.txt -massdns ~/massdns/bin/massdns -silent"
```

## Shuffledns Command for Docker

```
/root/go/bin/shuffledns -d {{rootDomain}} -r /app/Content/wordlists/dns_resolver_enum/default.txt -w /app/Content/wordlists/subdomain_enum/default.txt -massdns /app/massdns/bin/massdns -silent
```

## Shuffledns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Shuffledns/Script)

## Shuffledns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run shuffledns inside the docker

RUN apt-get update && apt-get install -y git wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN cd /app && wget https://raw.githubusercontent.com/reconness/reconness-agents/master/resolvers.txt
RUN wget https://github.com/projectdiscovery/shuffledns/releases/download/v1.0.4/shuffledns_1.0.4_linux_amd64.tar.gz
RUN tar -xzvf shuffledns_1.0.4_linux_amd64.tar.gz
RUN mv shuffledns /usr/local/bin/

# -------- End Agents dependencies -------- 
```
