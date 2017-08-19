#!/bin/sh

case "$(uname -s)" in

   Darwin)
     ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
     brew install go
     mkdir $HOME/go
     export GOPATH=$HOME/go
     export PATH=$PATH:$(go env GOPATH)/bin
     echo 'export GOPATH=$HOME/go' >>~/.bash_profile
     echo 'export PATH=$PATH:$GOPATH/bin' >>~/.bash_profile
     ;;

   Linux)
    sudo apt-get update
    sudo apt-get -y upgrade
    wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
    sudo tar -xvf go1.8.3.linux-amd64.tar.gz
    sudo mv go /usr/local
    echo 'export GOROOT=/usr/local/go' >>~/.profile
    echo 'export GOPATH=$HOME/go' >>~/.profile
    echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >>~/.profile
     ;;

   CYGWIN*|MINGW32*|MSYS*)
     echo 'Windows'
     ;;

   *)
     echo 'other OS'
     ;;
esac

go get github.com/nlack/ews-qr-app/tree/master/restapi
go get github.com/gamegos/jsend
go get github.com/go-playground/validator
go get github.com/knq/dburl
go get github.com/emicklei/go-restful
go get github.com/emicklei/go-restful-swagger12
go get github.com/joho/godotenv
cd $GOPATH/src/github.com/nlack/ews-qr-app/restapi
go run api.go #TODO build?
