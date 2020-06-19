OneForAll Dockerfile Entry

RUN git clone https://github.com/shmilylty/OneForAll.git
RUN cd OneForAll/
RUN python -m pip install -U pip setuptools wheel
RUN pip3 install -r requirements.txt
RUN python3 oneforall.py --help

## OneForAll Command

Using {{domain}} ReconNess replace {{domain}} for the subdomain.


```
./OneForAllWrapper.go -d {{domain}} 
```
## OneForAll Command for Docker

```
cd /app/OneForAllWrapper && go run OneForAllWrapper.go -d {{domain}} 
```

## OneForAll Script

Check [Script file](add script file location)

## OneForAll Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run gobuster inside the docker

RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
RUN mkdir OneForAllWrapper && cd OneForAllWrapper && wget https://raw.githubusercontent.com/hiddengearz/reconness-agents/master/OneForAll/OneForAllWrapper.go

# -------- End Agents dependencies -------- 
```


