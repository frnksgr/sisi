FROM golang:1.10.2
WORKDIR /go/src/github.com/frnksgr/sisi
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/github.com/frnksgr/sisi/sisi .
ENTRYPOINT ["/sisi", "-s"]