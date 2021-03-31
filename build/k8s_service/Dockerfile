FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/k8s
RUN cd /go/k8s/cmd/k8s_service && CGO_ENABLED=0 go build

FROM alpine:3.13.2
COPY --from=gobuilder /go/k8s/cmd/k8s_service/k8s_service /app/k8s_service
WORKDIR /app
ENTRYPOINT ["./k8s_service"]