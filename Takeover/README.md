## Takeover Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
python3 takeover.py -d {{domain}} -v
```

## Takeover Command for Docker

```
python3 /app/takeover/takeover.py -d {{domain}} -v
```

## Takeover Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Takeover/Script)

## Takeover Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run takeover inside the docker

RUN apt-get update && apt-get install -y git python3 python3-pip wget
RUN git clone https://github.com/m4ll0k/takeover.git
RUN cd takeover && python3 setup.py install

# -------- End Agents dependencies --------  
```
