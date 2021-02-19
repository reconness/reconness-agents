## Corsy Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.

```
python3 corsy.py -u https://{{domain}}
```

## Corsy Command for Docker
 
```
python3 /app/Corsy/corsy.py -u https://{{domain}}
```

## Corsy Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Corsy/Script)

## Corsy Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run corsy inside the docker

RUN apt-get update && apt-get install -y git python3 python3-pip
RUN git clone https://github.com/s0md3v/Corsy.git
RUN cd Corsy && pip3 install -r requirements.txt

# -------- End Agents dependencies -------- 
```
