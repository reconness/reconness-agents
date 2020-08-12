## Massdns command

This agent utilizes the [Reconness Unviersal Wrapper](https://github.com/hiddengearz/reconness-universal-wrapper) which must be installed for this agent to work. Using {{target}} ReconNess replaces {{{target}}} for the target and {{rootDomainName}} for the root domain. Ex: yahoo.com


If we have MassdnsWrapper in the folder ~/Desktop/MassdnsWrapper/

```
./reconness-universal-wrapper exec "/app/massdns/bin/massdns -r /app/massdns/lists/resolvers.txt *subdomains -w *outputFile -o S" -a api/targets/exportSubdomains/{{target}}/{{rootDomain}} --silent
```

## Massdns Command for Docker

```
/root/go/bin/reconness-universal-wrapper exec "/app/massdns/bin/massdns -r /app/massdns/lists/resolvers.txt *subdomains -w *outputFile -o S" -a api/targets/exportSubdomains/{{target}}/{{rootDomain}} --silent

```

## Massdns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Massdns/Script)


## Massdns Dockerfile Entry


# -------- Agents dependencies -------- 

```
# To allow run Massdns inside the docker

RUN apt-get update && apt-get install -y git build-essential wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN cd /app && /usr/local/go/bin/go get -u github.com/hiddengearz/reconness-universal-wrapper
RUN cd /root/go/bin/ && ./reconness-universal-wrapper setup -u <reconness username> -p <reconness password> -s <reconness.mydomain.com>
```


# -------- End Agents dependencies -------- 
