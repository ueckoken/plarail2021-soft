FROM golang:1.17-bullseye as build

WORKDIR /positioning
ADD . /positioning

RUN go build -o /positioning/app ./cmd/main.go 

FROM ubuntu:20.10

RUN apt-get update

COPY --from=build /positioning/app /app/app

CMD ["/app/app"]
