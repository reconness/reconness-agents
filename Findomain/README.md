## Findomain Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target.

```
./findomain-linux -t {{domain}} -r
```

## Findomain Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Findomain/Script)

## Findomain Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run findomain inside the docker

RUN apt-get update && apt-get install -y wget
RUN wget https://github.com/Edu4rdSHL/findomain/releases/latest/download/findomain-linux
RUN chmod +x findomain-linux

# -------- End Agents dependencies --------
```
