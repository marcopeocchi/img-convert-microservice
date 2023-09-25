FROM golang:bookworm AS build

RUN apt-get update
RUN apt-get install -y gcc libmupdf-dev binutils-dev libwebp-dev

WORKDIR /usr/src/fuku
COPY . .

RUN go get
RUN go build -o fuku *.go

FROM ubuntu

RUN apt-get update
RUN apt-get install -y libmupdf-dev libwebp-dev

WORKDIR /
COPY --from=build /usr/src/fuku /usr/bin

EXPOSE 8083

ENTRYPOINT [ "fuku", "-p", "8083" ]