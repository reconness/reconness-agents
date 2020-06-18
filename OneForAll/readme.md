OneForAll Dockerfile Entry

RUN git clone https://github.com/shmilylty/OneForAll.git
RUN cd OneForAll/
RUN python -m pip install -U pip setuptools wheel
RUN pip3 install -r requirements.txt
RUN python3 oneforall.py --help


