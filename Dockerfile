FROM golang:1.13 AS builder

WORKDIR $GOPATH/src/fiber-cloud-native

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM scratch

COPY --from=builder /go/src/fiber-cloud-native/main ./

EXPOSE 8080

ENTRYPOINT ["./main"]