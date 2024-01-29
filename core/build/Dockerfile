FROM ubuntu:latest
LABEL autor=mmmmaxonchik
# Setup ubuntu
RUN apt-get update
RUN apt-get -y install curl
# Setup go
RUN curl -OL https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
RUN rm -rf /usr/local/go 
RUN tar -C /usr/local -xvf go1.21.6.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
# Setup repository
RUN mkdir -p /home/core
COPY . /home/core
RUN cd /home/core
RUN cd /home/core && go build /home/core/cmd/server/main.go 
ENV CONFIG_PATH="/home/core/config/local.yaml"
ENTRYPOINT [ "/home/core/main" ]