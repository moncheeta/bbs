FROM golang:1.22.3-alpine
RUN apk add git cmake build-base zlib-dev openssl-dev libuv-dev json-c-dev libwebsockets-dev libwebsockets-evlib_uv

WORKDIR /src/ttyd
COPY ./ttyd/* ./
RUN mkdir ./build && cd ./build && cmake .. && make && make install

WORKDIR /src/bbs
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./client
COPY ./assets ./assets

EXPOSE 80
CMD ["ttyd", "-p", "80", "-W", "./client"]
