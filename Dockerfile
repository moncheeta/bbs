FROM golang:1.22.3-alpine
RUN apk add cmake build-base zlib-dev openssl-dev libuv-dev json-c-dev libwebsockets-dev libwebsockets-evlib_uv

WORKDIR /src
COPY ./ttyd ./ttyd
RUN mkdir -p ./ttyd/build && cd ./ttyd/build && rm -r ./* && cmake .. && make && make install

WORKDIR /src/bbs
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY ./assets ./assets
RUN CGO_ENABLED=0 GOOS=linux go build -o ./client

EXPOSE 80
CMD ["ttyd", "-p", "443", "-t", "titleFixed=Damian's BBS", "-t", "disableLeaveAlert=true", "-t", "disableResizeOverlay=true", "-S", "-C", "/ssl/fullchain.pem", "-K", "/ssl/privkey.pem", "-W", "./client"]
