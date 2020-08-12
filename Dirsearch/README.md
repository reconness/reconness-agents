## Dirsearch Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
python3 dirsearch.py -u {{domain}} -e php -r
```
## Dirsearch Command for Docker

```
/app/dirsearch/dirsearch.py -u {{domain}} -e php -r
```

## Dirsearch Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Dirsearch/Script)

## Dirsearch Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run dirsearch inside the docker

RUN apt-get update && apt-get install -y git python3 python3-pip
RUN git clone https://github.com/maurosoria/dirsearch.git

# -------- End Agents dependencies -------- 
```
