## Puredns command

Using {{rootDomain}} ReconNess replace {{rootDomain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target

If we have massdns in the folder /app/massdns/bin/massdns 

```
puredns bruteforce subdomains.txt {{rootDomain}} --resolvers resolvers.txt --bin /app/massdns/bin/massdns -q
```

## Puredns Command for Docker

```
/root/go/bin/puredns bruteforce /app/Content/wordlists/subdomain_enum/default.txt {{rootDomain}} --resolvers /app/Content/wordlists/dns_resolver_enum/default.txt --bin /app/massdns/bin/massdns -q

```

## Puredns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Puredns/Script)

## Puredns Dockerfile Entry

# -------- Agents dependencies -------- 

```
# To allow run Puredns inside the docker

RUN apt-get update && apt-get install -y git build-essential wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN GO111MODULE=on /usr/local/go/bin/go get github.com/d3mondev/puredns/v2
```

# -------- End Agents dependencies -------- 
