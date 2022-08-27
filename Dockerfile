FROM golang:alpine AS builder
RUN apk --update --no-cache add build-base git
WORKDIR $GOPATH/src/git.mnc.sh/ilazarev/trade
COPY . ./
RUN go get
RUN go generate ./...
RUN go test ./...
RUN go build -ldflags "-X main.Version=$(git describe --tags --long --always)" ./cmd/trade
RUN cp trade /bin/trade

FROM alpine:latest
RUN apk --update --no-cache add ca-certificates icu-libs
COPY --from=builder /bin/trade /bin/trade
WORKDIR /data
EXPOSE 3000/tcp
ENTRYPOINT ["/bin/trade"]
