# -------- Agents dependencies -------- 

# To allow run Massdns inside the docker

RUN apt-get update && apt-get install -y git python3 python3-pip
RUN git clone https://github.com/nsonaniya2010/SubDomainizer.git
RUN cd SubDomainizer && pip3 install -r requirements.txt

# -------- End Agents dependencies -------- 
```