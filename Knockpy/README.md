## Knockpy Command

Using {{domain}} ReconNess replace {{domain}} to the root domain or {{rootDomain}}, for example, yahoo.com if we define that as a root domain adding the Target.

If we have knockpy in the folder /app/knock/

```
python3 /app/knock/knockpy/knockpy.py {{rootDomain}}
```

## Knockpy Command for Docker
 
```
python3 /app/knock/knockpy/knockpy.py {{rootDomain}}
```

## Knockpy Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Knockpy/Script)

## Knockpy Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run knockpy the docker

RUN apt-get update && apt-get install -y git python2.7 python-pip python-dnspython
RUN git clone https://github.com/guelfoweb/knock
RUN cd knock && pip3 install -r requirements.txt

# -------- End Agents dependencies -------- 
```
