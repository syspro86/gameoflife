FROM  golang:1.16 as builder
WORKDIR /tmp/tiny-golang-image
COPY . .
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
RUN GOOS=js GOARCH=wasm go build -o main.wasm cmd/gol/main.go
