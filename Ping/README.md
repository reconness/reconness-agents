## Ping Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
ping {{domain}} -c 1
```

## Ping Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Ping/Script)

## Ping Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run ping inside the docker

RUN apt-get update && apt-get install -y iputils-ping

# -------- End Agents dependencies -------- 
```
