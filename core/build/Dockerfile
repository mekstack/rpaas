FROM golang:bookworm
LABEL autor=mmmmaxonchik
# Setup repository
RUN mkdir /project
COPY . /project
RUN ls /project
RUN cd /project 
RUN cd /project && go build ./cmd/server/main.go 
# Setup core
ENV NATAAS_REDIS_ADDR="redis:6379"
ENV NATAAS_GRPC_ADDR="127.0.0.1:8080"
ENV NATAAS_ENVIRONMENT=0

ENTRYPOINT [ "/project/main" ]
