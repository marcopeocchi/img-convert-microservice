FROM golang:bookworm AS build

RUN apt-get update && apt-get install gcc libvips-dev libavif-dev libheif-dev -y

WORKDIR /usr/src/fuku
COPY . .

RUN go get
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o fuku main.go

FROM ubuntu

RUN apt-get update && apt-get install libvips-dev libavif-dev libheif-dev -y

WORKDIR /
COPY --from=build /usr/src/fuku /usr/bin

EXPOSE 8083

ENTRYPOINT [ "fuku", "-p", "8083" ]