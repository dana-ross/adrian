FROM alpine:latest

EXPOSE 80/tcp 443/tcp

RUN wget https://github.com/daveross/adrian/releases/download/v`wget -q -O - https://api.github.com/repos/daveross/adrian/releases/latest | grep tag_name | tr -s ' ' | cut -d ' ' -f3 | sed s/[^0-9\.]//g`/adrian_`wget -q -O - https://api.github.com/repos/daveross/adrian/releases/latest | grep tag_name | tr -s ' ' | cut -d ' ' -f3 | sed s/[^0-9\.]//g`_`uname -s`_`uname -m`.tar.gz

RUN tar xvfz adrian*.tar.gz

RUN rm *.tar.gz

RUN mkdir -p /fonts

ADD adrian.yaml.docker.example /etc/adrian.yaml

CMD adrian/adrian --config /etc/adrian.yaml