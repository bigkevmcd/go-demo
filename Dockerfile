FROM golang:latest AS build
WORKDIR /go/src
COPY . /go/src
RUN go build ./cmd/go-demo

FROM gcr.io/distroless/base-debian11
WORKDIR /root/
COPY --from=build /go/src/go-demo .
EXPOSE 8080
ENTRYPOINT ["./go-demo"]
