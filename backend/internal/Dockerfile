FROM golang:1.17-bullseye as build

WORKDIR /control
ADD . /control

RUN go build -o /control/app ./cmd/main.go 


FROM ubuntu:20.10

COPY --from=build /control/app /app/app

CMD ["/app/app"]
