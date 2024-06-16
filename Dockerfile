FROM golang:alpine as build
WORKDIR /usr/src

COPY ./ttyd ./ttyd
RUN apk add --no-cache cmake build-base zlib-dev openssl-dev libuv-dev json-c-dev libwebsockets-dev libwebsockets-evlib_uv
WORKDIR /usr/src/ttyd
RUN mkdir ./build && cd ./build && cmake .. && make

WORKDIR /usr/src/bbs
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY ./assets ./assets
RUN CGO_ENABLED=0 GOOS=linux go build -o ./client

FROM alpine
RUN apk add --no-cache zlib-dev openssl-dev libuv-dev json-c-dev libwebsockets-dev libwebsockets-evlib_uv
COPY --from=build /usr/src/ttyd/build/ttyd /bin/ttyd
COPY --from=build /usr/src/bbs/client /bin/client
COPY --from=build /usr/src/bbs/assets/ /assets/
EXPOSE 443
CMD ["ttyd", "-p", "443", "-t", "titleFixed=Damian's BBS", "-t", "disableLeaveAlert=true", "-t", "disableResizeOverlay=true", "-S", "-C", "/ssl/fullchain.pem", "-K", "/ssl/privkey.pem", "-W", "client"]
