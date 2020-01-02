## Nmap Command

```
nmap
```

## Nmap Arguments

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target.

```
-T4 {{domain}}
```

## Nmap Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Nmap/Script)

## Nmap Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run nmap inside the docker
$ RUN apt-get update && apt-get install -y nmap

# -------- End Agents dependencies -------- 
```
