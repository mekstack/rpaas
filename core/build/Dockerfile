FROM golang:bookworm
LABEL autor=mmmmaxonchik
# Setup repository
RUN mkdir /project
WORKDIR /project
COPY . .
RUN go build ./cmd/server/main.go 
ENTRYPOINT [ "/project/main" ]
