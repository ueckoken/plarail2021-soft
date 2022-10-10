FROM golang:1.17-bullseye AS builder

ENV CGO_ENABLED=0
WORKDIR /control
COPY ./go.mod ./go.sum /control/
RUN go mod download
COPY . /control
RUN make build
# RUN go build -o /control/app ./cmd/main.go 

FROM gcr.io/distroless/static-debian11:nonroot AS runner
COPY --from=builder --chown=nonroot:nonroot /control/external /external
ENTRYPOINT [ "/external" ]
