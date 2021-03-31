FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/k8s
RUN cd /go/k8s/cmd/watch_service && CGO_ENABLED=0 go build

FROM alpine:3.13.2
COPY --from=gobuilder /go/k8s/cmd/watch_service/watch_service /app/watch_service
WORKDIR /app
ENTRYPOINT ["./watch_service"]