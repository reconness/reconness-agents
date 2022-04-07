FROM mcr.microsoft.com/dotnet/sdk:6.0.100-bullseye-slim AS build
WORKDIR /app

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs
RUN npm install -g @vue/cli

# copy csproj and restore as distinct layers
COPY *.sln .
COPY ["DAL/ReconNess.Data.Npgsql/ReconNess.Data.Npgsql.csproj", "DAL/ReconNess.Data.Npgsql/"]
COPY ["ReconNess.Web/ReconNess.Web.csproj", "ReconNess.Web/"]
COPY ["ReconNess.Entities/ReconNess.Entities.csproj", "ReconNess.Entities/"]
COPY ["ReconNess.Core/ReconNess.Core.csproj", "ReconNess.Core/"]
COPY ["ReconNess.Worker/ReconNess.Worker.csproj", "ReconNess.Worker/"]
COPY ["ReconNess/ReconNess.csproj", "ReconNess/"]
RUN dotnet restore "ReconNess.Web/ReconNess.Web.csproj"

# copy everything else and build app
COPY . ./
WORKDIR /app/ReconNess.Web
RUN dotnet publish -c Release -o /dist

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS runtime
WORKDIR /app

#####################################################################################################################
# If you want to generate your own certificate with different password
# you can run
#
# dotnet dev-certs https -ep %USERPROFILE%\.aspnet\https\aspnetapp.pfx -p { password here }
# dotnet dev-certs https --trust
# 
# and replace `reconness\src\aspnetapp.pfx` with the file `%USERPROFILE%\.aspnet\https\aspnetapp.pfx` generated
# and replace the password that you used `{ password here }` 
# ENV ASPNETCORE_Kestrel__Certificates__Default__Password="{ password here }"
#####################################################################################################################

COPY aspnetapp.pfx .

ENV ASPNETCORE_URLS http://+:5000;https://+:5001
ENV ASPNETCORE_Kestrel__Certificates__Default__Password="password"
ENV ASPNETCORE_Kestrel__Certificates__Default__Path="aspnetapp.pfx"
EXPOSE 5000
EXPOSE 5001

COPY --from=build /dist ./

# -------- Agents dependencies -------- 

# To allow run the wrapper 
################################################################################################################################################
# Change <reconness username> <reconness password> <reconness.mydomain.com> with your username, password and domain where reconness is running.
#
# Ex.
#
# RUN /root/go/bin/reconness-universal-wrapper setup -u myusername -p mypasssord -s http://mydomainorip.com
#################################################################################################################################################
RUN apt-get update && apt-get install -y git wget unzip python2.7 python-pip python3 python3-pip python-dnspython build-essential
# Install Golang
RUN wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Enable Go modules support
ENV GO111MODULE=on

RUN /usr/local/go/bin/go get -u github.com/hiddengearz/reconness-universal-wrapper
RUN /root/go/bin/reconness-universal-wrapper setup -u <reconness username> -p <reconness password> -s <reconness.mydomain.com>

# To allow run subfinder inside the docker
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

# To allow run amass inside the docker
RUN cd /tmp/ ; wget https://github.com/OWASP/Amass/releases/download/v3.7.4/amass_linux_amd64.zip ; unzip amass_linux_amd64.zip
RUN mv /tmp/amass_linux_amd64/amass /bin

# To allow run gobuster inside the docker
RUN /usr/local/go/bin/go install github.com/OJ/gobuster/v3@latest

# To allow run sublist3r inside the docker
RUN git clone https://github.com/aboul3la/Sublist3r.git
RUN cd Sublist3r && pip install -r requirements.txt

# To allow run findomain inside the docker
RUN wget https://github.com/Edu4rdSHL/findomain/releases/latest/download/findomain-linux
RUN chmod +x findomain-linux

# To allow run ffuf inside the docker
RUN /usr/local/go/bin/go install github.com/ffuf/ffuf@latest

# To allow run httprobe inside the docker
RUN /usr/local/go/bin/go get github.com/tomnomnom/httprobe

# To allow run ping inside the docker
RUN apt-get install -y iputils-ping

# To allow run nmap inside the docker
RUN apt-get install -y nmap

# To allow run takeover inside the docker
RUN git clone https://github.com/m4ll0k/takeover.git
RUN cd takeover && python3 setup.py install

# To allow run OneForAll inside the docker
RUN git clone https://github.com/shmilylty/OneForAll.git
RUN python3 -m pip install -U pip setuptools wheel
RUN pip3 install -r /app/OneForAll/requirements.txt

# To allow run zdns inside the docker
RUN git clone https://github.com/zmap/zdns.git
RUN cd zdns && /usr/local/go/bin/go build

# To allow run knockpy the docker
RUN git clone https://github.com/guelfoweb/knock
RUN cd knock && pip3 install -r requirements.txt

# To allow run Massdns inside the docker
RUN git clone https://github.com/blechschmidt/massdns.git && cd massdns && make

# To allow run waybackurls inside the docker
RUN /usr/local/go/bin/go install github.com/tomnomnom/waybackurls@latest

# To allow run gau inside the docker
RUN /usr/local/go/bin/go install github.com/lc/gau/v2/cmd/gau@latest

# To allow run naabu inside the docker
RUN wget https://github.com/projectdiscovery/naabu/releases/download/v2.0.3/naabu-linux-amd64.tar.gz
RUN tar -xvf naabu-linux-amd64.tar.gz
RUN cp naabu-linux-amd64 /usr/local/bin/naabu

# To allow run shuffledns inside the docker
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/shuffledns/cmd/shuffledns@latest

# To allow run corsy inside the docker
RUN git clone https://github.com/s0md3v/Corsy.git
RUN cd Corsy && pip3 install -r requirements.txt

# To allow run dirsearch inside the docker
RUN git clone https://github.com/maurosoria/dirsearch.git
RUN cd dirsearch && pip3 install -r requirements.txt

# To allow run dnsx inside the docker
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/dnsx/cmd/dnsx@latest

# To allow run nuclei inside the docker
RUN /usr/local/go/bin/go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
RUN /root/go/bin/nuclei -update-templates

# To allow run crlfuzz inside the docker
RUN /usr/local/go/bin/go install -v github.com/dwisiswant0/crlfuzz/cmd/crlfuzz@latest

# To allow run puredns inside the docker
RUN /usr/local/go/bin/go install -v github.com/d3mondev/puredns/v2@latest

# To allow run gowitness inside the docker
RUN /usr/local/go/bin/go install -v github.com/sensepost/gowitness@latest

# -------- End Agents dependencies -------- 

ENTRYPOINT ["dotnet", "ReconNess.Web.dll"]
