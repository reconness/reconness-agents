## OneForAll Command

This agent utilizes the [Reconness Unviersal Wrapper](https://github.com/hiddengearz/reconness-universal-wrapper) which must be installed for this agent to work. Using {{rootDomain}} ReconNess replace {{rootDomain}} for the root domain. Ex: yahoo.com

If we have OneForAllWrapper in the folder ~/Desktop/OneForAllWrapper/

```
cd ~/Desktop/OneForAllWrapper/ && ./reconness-universal-wrapper exec "python3 /app/OneForAll/oneforall.py --target {{rootDomain}} -path *outputDir/*.txt run" --silent
```

## OneForAll Command for Docker

```
/root/go/bin/reconness-universal-wrapper exec "python3 /app/OneForAll/oneforall.py --target {{rootDomain}} -path *outputDir/*.txt run" --silent
```

## OneForAll Script

Check [Script file](add script file location)

## OneForAll Dockerfile Entry


# -------- Agents dependencies -------- 

# To allow run OneForAll inside the docker

```
RUN apt-get update && apt-get install -y git wget python3 python3-pip
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN git clone https://github.com/shmilylty/OneForAll.git
RUN python3 -m pip install -U pip setuptools wheel
RUN pip3 install -r /app/OneForAll/requirements.txt
RUN /usr/local/go/bin/go get -u github.com/hiddengearz/reconness-universal-wrapper
RUN /root/go/bin/reconness-universal-wrapper setup -u <reconness username> -p <reconness password> -s <reconness.mydomain.com>
```
# -------- End Agents dependencies -------- 
