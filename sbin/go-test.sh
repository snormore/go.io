#!/bin/bash
CONFIG_PATH="`pwd`/config"
CODE_DIR="`pwd`/gonode/src/go.io"
cd $CODE_DIR
for pkg in $(find ./*/ -type d)
do 
    echo "Running '${pkg:2}' tests..."
    cd "$CODE_DIR/$pkg"
    export CONFIG_PATH=$CONFIG_PATH
    export GO_ENV="test"
    go test -i
    go test
done
