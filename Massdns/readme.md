## Massdns command

Using {{target}} ReconNess replaces {{target}} for the target and {{rootDomain}} for the root domain. Ex: yahoo.com

If we have MassdnsWrapper in the folder ~/Desktop/MassdnsWrapper/

```
cd ~/Desktop/MassdnsWrapper && /usr/local/go/bin/go run MassdnsWrapper.go -o "-t A --ignore --norecurse --predictable" -b http://localhost -u {{userName}} -p {{password}} -s api/targets/exportSubdomains/{{target}}/{{rootDomain}}
```

## Massdns Command for Docker

```
/usr/local/go/bin/go run MassdnsWrapper.go -o "-t A --ignore --norecurse --predictable" -b http://localhost -u {{userName}} -p {{password}} -s api/targets/exportSubdomains/{{target}}/{{rootDomain}}
```

## Massdns Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Massdns/Script)


## Massdns Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run Massdns inside the docker

RUN apt-get update && apt-get install -y git build-essential wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make
RUN cd /app && wget https://raw.githubusercontent.com/reconness/reconness-agents/master/Massdns/MassdnsWrapper.go

# -------- End Agents dependencies -------- 
```
