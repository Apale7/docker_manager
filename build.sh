#!/usr/bin/env bash
NAME="docker_manager"
#编译pb
# protoc --go_out=. --go_opt=paths=source_relative \
#     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#     proto/*/*.proto
find ./ -name "*.go" | xargs gofmt -w -s -l
mkdir -p output/bin output/config
cp -r config/ output/


name=`uname -o`
if [[ $name =~ "Darwin" ]];then
    GOOS=darwin GOARCH=amd64 go build -o output/bin/${NAME}.out
    chmod +x output/bin/${NAME}.out
elif [[ $name =~ "GNU/Linux" ]];then
    GOOS=linux GOARCH=amd64 go build -o output/bin/${NAME}.out
    chmod +x output/bin/${NAME}.out
else
    GOOS=windows GOARCH=amd64 go build -o output/bin/${NAME}.exe
fi
chmod +x output/start.sh