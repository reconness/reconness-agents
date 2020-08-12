## OneForAll Command

Using {{rootDomain}} ReconNess replace {{rootDomain}} for the root domain. Ex: yahoo.com

If we have OneForAllWrapper in the folder ~/Desktop/OneForAllWrapper/

```
cd ~/Desktop/OneForAllWrapper && /usr/local/go/bin/go run OneForAllWrapper.go -d {{rootDomain}} 
```

## OneForAll Command for Docker

```
/usr/local/go/bin/go run OneForAllWrapper.go -d {{rootDomain}}
```

## OneForAll Script

Check [Script file](add script file location)

## OneForAll Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run OneForAll inside the docker

RUN apt-get update && apt-get install -y git python3 python3-pip wget
RUN wget https://dl.google.com/go/go1.14.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
RUN wget https://raw.githubusercontent.com/reconness/reconness-agents/master/OneForAll/OneForAllWrapper.go
RUN git clone https://github.com/shmilylty/OneForAll.git
RUN python3 -m pip install -U pip setuptools wheel
RUN pip3 install -r /app/OneForAll/requirements.txt

# -------- End Agents dependencies -------- 
```
