#!/bin/bash

function gen_grpc() {
  protoc -I cmd/watch_service/protos watcher.proto --go_out=cmd/watch_service --go-grpc_out=cmd/watch_service
}

function build_image() {
    docker build -t lqha.xyz/k8s-test:latest .
}

function print_usage() {
    echo "usage: ./build.sh [gen|build]"
    exit
}

cmd=$1

case $cmd in
  "gen")
  gen_grpc
  ;;

  "build")
  build_image
  ;;

  *)
  print_usage
  ;;
esac



