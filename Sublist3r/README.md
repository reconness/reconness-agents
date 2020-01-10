## Sublist3r Command

Using {{domain}} ReconNess replace {{domain}} to the root domain, for example, yahoo.com if we define that as a root domain adding the Target.

If we have sublist3r in the folder /app/Sublist3r/

```
python /app/Sublist3r/sublist3r.py -d {{domain}}
```

## Sublist3r Script

Check [Script file](https://github.com/reconness/reconness-agents/blob/master/Sublist3r/Script)

## Sublist3r Dockerfile Entry

```
# -------- Agents dependencies -------- 

# To allow run sublist3r inside the docker
$ RUN apt-get install -y git
$ RUN apt-get install -y python2.7 python-pip
$ RUN git clone https://github.com/aboul3la/Sublist3r.git
$ RUN cd Sublist3r && pip install -r requirements.txt

# -------- End Agents dependencies --------
```
