#!/bin/bash

APP=$1
VERSION=$2

#targets OS, you can add more to this OS collections
PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64")

if [[ -z "$APP" || -z "$VERSION" ]]; then
    echo "$0 require package name and version arguments"
    echo "example : ./build.sh github.com/Bhinneka/kece/cmd v1.0.0"
    exit 1
fi

#split APP variable
APP_SPLIT=(${APP//\// })

#access app name using 2nd index
APP_NAME=${APP_SPLIT[2]}

for PLATFORM in "${PLATFORMS[@]}"
do
    PLATFORM_SPLIT=(${PLATFORM//\// })
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH=${PLATFORM_SPLIT[1]}
    OUTPUT_BINARY=$APP_NAME
    if [ $GOOS = "windows" ]; then
        OUTPUT_BINARY+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_BINARY $APP
    #when something goes wrong with go build, just exit imidiately
    if [ $? -ne 0 ]; then
        echo 'go build fail...'
        exit 1
    fi

    #OUTPUT_TAR_NAME=$APP_NAME'-v0.0.0.'$GOOS'-'$GOARCH'.tar.gz'
    OUTPUT_TAR_NAME=$APP_NAME'-'$VERSION'.'$GOOS'-'$GOARCH'.tar.gz'
    env tar czf $OUTPUT_TAR_NAME $OUTPUT_BINARY
    #when something goes wrong with tar, just exit imidiately
    if [ $? -ne 0 ]; then
        echo 'tar process fail...'
        exit 1
    fi
    rm $OUTPUT_BINARY
done

