#!/bin/bash

function gen_grpc() {
  cd ../
  protoc -I cmd/watch_service/protos watcher.proto --go_out=cmd/watch_service --go-grpc_out=cmd/watch_service
}

function build_image() {
  service=$1
  cur_dir=$(pwd)
  root_dir=${cur_dir}/../
  cd ${root_dir} || exit

  if [[ "$service" != "k8s_service" && "$service" != "watch_service" ]];then
    print_usage
  fi
  cp ${cur_dir}/$service/Dockerfile ${root_dir}/Dockerfile && docker build -t lqha.xyz/$service:latest . && rm -rf ${root_dir}/Dockerfile
}

function print_usage() {
  echo "usage: ./build.sh [gen|build] [k8s_service|watch_service]"
  exit
}

cmd=$1
service=$2

case $cmd in
  "gen")
  gen_grpc
  ;;

  "build")
  build_image $service
  ;;

  *)
  print_usage
  ;;
esac

