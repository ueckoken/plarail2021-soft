FROM golang:1.17-bullseye as builder

WORKDIR /speed
ADD . /speed

RUN go build -o /speed/app ./cmd/main.go


FROM ubuntu:20.10 as runner

COPY --from=builder /speed/app /app/app

CMD ["/app/app"]