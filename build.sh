#!/bin/bash
docker build -f docker/Dockerfile -t gol_build .
docker create --name gol_build gol_build
docker cp gol_build:/tmp/tiny-golang-image/main.wasm .
docker cp gol_build:/tmp/tiny-golang-image/wasm_exec.js .
docker rm gol_build
