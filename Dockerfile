FROM golang:alpine
WORKDIR /app
ADD . /app
RUN go build -o goapp
USER 1001
ENTRYPOINT [ "/app/goapp" ]