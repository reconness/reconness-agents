FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /app

RUN curl -sL https://deb.nodesource.com/setup_10.x |  bash -
RUN apt-get install -y nodejs
RUN npm install -g @vue/cli

# copy csproj and restore as distinct layers
COPY *.sln .
COPY ["DAL/ReconNess.Data.Npgsql/ReconNess.Data.Npgsql.csproj", "DAL/ReconNess.Data.Npgsql/"]
COPY ["ReconNess.Web/ReconNess.Web.csproj", "ReconNess.Web/"]
COPY ["ReconNess.Entities/ReconNess.Entities.csproj", "ReconNess.Entities/"]
COPY ["ReconNess.Core/ReconNess.Core.csproj", "ReconNess.Core/"]
COPY ["ReconNess/ReconNess.csproj", "ReconNess/"]
RUN dotnet restore "ReconNess.Web/ReconNess.Web.csproj"

# copy everything else and build app
COPY . ./
WORKDIR /app/ReconNess.Web
RUN dotnet publish -c Release -o /dist

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1 AS runtime
WORKDIR /app

ENV ASPNETCORE_URLS http://+:5000
EXPOSE 5000

COPY --from=build /dist ./

# -------- Agents dependencies -------- 

# To allow run subfinder inside the docker
RUN apt-get update && apt-get install -y git wget unzip
RUN wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
RUN echo 'export GOROOT=/usr/local/go' >> ~/.profile
RUN echo 'export GOPATH=$HOME/go'	>> ~/.profile
RUN echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.profile
RUN . ~/.profile && go get -u github.com/projectdiscovery/subfinder/cmd/subfinder

# To allow run amass inside the docker
RUN cd /tmp/ ; wget https://github.com/OWASP/Amass/releases/download/v3.7.4/amass_linux_amd64.zip ; unzip amass_linux_amd64.zip
RUN mv /tmp/amass_linux_amd64.zip/amass /bin

# To allow run gobuster inside the docker
RUN . ~/.profile && go get github.com/OJ/gobuster
RUN wget https://gist.githubusercontent.com/gorums/e471b842f1dd633469bffb5f750e0325/raw/96f4e51d96b2203f19f6381c8c545b278eaa0837/all.txt

# To allow run sublist3r inside the docker
RUN apt-get install -y python2.7 python-pip
RUN git clone https://github.com/aboul3la/Sublist3r.git
RUN cd Sublist3r && pip install -r requirements.txt

# To allow run findomain inside the docker
RUN wget https://github.com/Edu4rdSHL/findomain/releases/latest/download/findomain-linux
RUN chmod +x findomain-linux

# To allow run ffuf inside the docker
RUN . ~/.profile && go get github.com/ffuf/ffuf
RUN wget https://gist.githubusercontent.com/gorums/0a3a9d903e8e47fbff9d91097e19b4f8/raw/c81a34fe84731430741e0463eb6076129c20c4c0/content_discovery_all.txt

# To allow run dirsearch inside the docker
RUN apt-get install -y python3 python3-pip
RUN git clone https://github.com/maurosoria/dirsearch.git

# To allow run httprobe inside the docker
RUN . ~/.profile && go get github.com/tomnomnom/httprobe

# To allow run ping inside the docker
RUN apt-get update && apt-get install -y iputils-ping

# To allow run nmap inside the docker
RUN apt-get update && apt-get install -y nmap

# To allow run takeover inside the docker
RUN git clone https://github.com/m4ll0k/takeover.git
RUN cd takeover && python3 setup.py install

# To allow run OneForAll inside the docker
RUN wget https://raw.githubusercontent.com/reconness/reconness-agents/master/OneForAll/OneForAllWrapper.go
RUN git clone https://github.com/shmilylty/OneForAll.git
RUN python3 -m pip install -U pip setuptools wheel
RUN pip3 install -r /app/OneForAll/requirements.txt

# To allow run zdns inside the docker
RUN git clone https://github.com/zmap/zdns.git
RUN . ~/.profile && cd zdns && go build

# To allow run knockpy the docker
RUN apt-get install -y python-dnspython
RUN git clone https://github.com/guelfoweb/knock
RUN cd knock && python setup.py install

# -------- End Agents dependencies -------- 

ENTRYPOINT ["dotnet", "ReconNess.Web.dll"]
