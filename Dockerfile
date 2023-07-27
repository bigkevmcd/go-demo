ARG TARGETOS
ARG TARGETARCH
ARG VERSION

FROM golang:latest AS build
WORKDIR /go/src
COPY . /go/src
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build ./cmd/go-demo

FROM gcr.io/distroless/base-debian11
WORKDIR /root/
COPY --from=build /go/src/go-demo .
EXPOSE 8080
ENTRYPOINT ["./go-demo"]
