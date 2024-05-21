FROM golang:alpine AS build

RUN apk update && apk add gcc musl-dev vips-dev poppler-dev libavif-dev libheif-dev

WORKDIR /usr/src/fuku
COPY . .

RUN go get
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o fuku main.go

FROM alpine

RUN apk update && apk add vips-dev poppler-dev libavif-dev libheif-dev --no-cache

WORKDIR /
COPY --from=build /usr/src/fuku /usr/bin

EXPOSE 8080

ENTRYPOINT [ "fuku", "-p", "8080" ]