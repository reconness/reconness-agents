## Naabu Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
./naabu -host {{domain}} -silent
```

## Naabu for Docker

```
/app/naabu -host {{domain}} -silent
```

## Naabu Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Naabu/Script)

## Naabu Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run naabu inside the docker

RUN wget https://github.com/projectdiscovery/naabu/releases/download/v2.0.2/naabu_2.0.2_linux_amd64.tar.gz
RUN tar -xvf naabu_2.0.2_linux_amd64.tar.gz

# -------- End Agents dependencies -------- 
```
