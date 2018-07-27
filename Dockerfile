FROM golang:alpine
WORKDIR /app
ADD . /app
RUN apk add --no-cache git && \
    go get -v -d && \
    go build -o goapp && \
    apk del git
USER 1001
ENTRYPOINT [ "/app/goapp" ]