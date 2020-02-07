## Nmap Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
nmap -T4 {{domain}}
```

## Nmap Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Nmap/Script)

## Nmap Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run nmap inside the docker
RUN apt-get update && apt-get install -y nmap

# -------- End Agents dependencies -------- 
```
