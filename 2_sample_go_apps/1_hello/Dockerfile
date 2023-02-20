FROM golang:latest as build
COPY main.go ./
RUN go build --ldflags "-linkmode=external" main.go

FROM scratch
ADD libs.tar.gz /
WORKDIR /go/
COPY --from=build ./go/main /go/
USER 2000
CMD ["./main"]

