# -------- Agents dependencies -------- 

# To allow run Massdns inside the docker

RUN apt-get update && apt-get install -y git
RUN apt-get install -y wget
RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make

# -------- End Agents dependencies -------- 
```