#!/bin/bash

install_kece() 
{
    OS=$1
    echo "install kece..."

    #check os from args
    case $OS in
        darwin)
            echo "os darwin"
            # TODO
            # wget https://github.com/Bhinneka/kece/releases/download/v0.0.0/kece-v0.0.0.darwin-amd64.tar.gz
            # tar -zxvf kece-v0.0.0.darwin-amd64.tar.gz
            ;;
        linux)
            echo "os linux"
            # TODO
            # wget https://github.com/Bhinneka/kece/releases/download/v0.0.0/kece-v0.0.0.linux-amd64.tar.gz
            # tar -zxvf kece-v0.0.0.linux-amd64.tar.gz
            ;;
        *)
            echo "operating system unknown"
            ;;
    esac
}

install_kece "$@"

# TODO
# curl https://raw.githubusercontent.com/Bhinneka/kece/master/scripts/install.sh | bash -s {your os name}