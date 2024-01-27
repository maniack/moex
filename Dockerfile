FROM golang:alpine AS builder
RUN apk --update --no-cache add build-base git
WORKDIR $GOPATH/src/github.com/maniack/moex_exporter
COPY . ./
RUN go get
RUN go generate ./...
RUN go test ./...
RUN go build -ldflags "-X main.Version=$(git describe --tags --long --always)" ./cmd/moex_exporter
RUN cp moex_exporter /bin/moex_exporter

FROM alpine:3.18.6
RUN apk --update --no-cache add ca-certificates
COPY --from=builder /bin/moex_exporter /bin/moex_exporter
EXPOSE 9927/tcp
ENTRYPOINT ["/bin/moex_exporter"]
