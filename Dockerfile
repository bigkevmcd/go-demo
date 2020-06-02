FROM golang:latest AS build
WORKDIR /go/src
COPY . /go/src
RUN go build ./cmd/go-demo

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /root/
COPY --from=build /go/src/go-demo .
EXPOSE 8080
CMD ["./go-demo"]
