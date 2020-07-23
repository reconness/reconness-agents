## Massdns command

Using {{target}} ReconNess replaces {{{target}}} for the target and {{rootDomainName}} for the root domain. Ex: yahoo.com

If we have MassdnsWrapper in the folder ~/Desktop/MassdnsWrapper/

```
cd ~/Desktop/MassdnsWrapper && usr/local/go run MassdnsWrapper.go -b https://localhost -a api/Auth/Login -u <username> -p <password> -s api/targets/exportSubdomains/{{target}}/{{rootDomain}}
```

## MassDNS Command for Docker

```
/usr/local/go/bin/go run MassdnsWrapper.go -b https://localhost -a api/Auth/Login -u <username> -p <password> -s api/targets/exportSubdomains/{{target}}/{{rootDomain}}
```

## Massdns Dockerfile Entry

# -------- Agents dependencies -------- 

# To allow run Massdns inside the docker

```
RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN cd /app && wget https://raw.githubusercontent.com/hiddengearz/reconness-agents/master/Massdns/MassdnsWrapper.go
```

# -------- End Agents dependencies -------- 
```
