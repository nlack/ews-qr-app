# REST API  
Backend with CRUD Operations

## BACKEND INSTALLATION

### System Requirements (Mac)
- brew
`ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
- mysql
`brew install mysql`
- golang
`brew install go`
### Configure Go
- set GOPATH
```
mkdir $HOME/go
export GOPATH=$HOME/go
export PATH=$PATH:$(go env GOPATH)/bin
echo 'export GOPATH=$HOME/go' >>~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin' >>~/.bash_profile
```
- get dependencies
```
go get github.com/nlack/ews-qr-app/restapi
go get github.com/gamegos/jsend
go get github.com/go-playground/validator
go get github.com/knq/dburl
go get github.com/emicklei/go-restful
go get github.com/emicklei/go-restful-swagger12
go get github.com/joho/godotenv
go get github.com/go-errors/errors
```
### Rename .env
```
cd $GOPATH/src/github.com/nlack/ews-qr-app/restapi
mv .env-sample .env
```

### Configure MySQL
- start service
```
brew tap homebrew/services
brew services start mysql
```
- set root password and login
```
mysqladmin -u root password 'ews'
mysql -u root -p
```
- create database and user
```
CREATE DATABASE ewsdb
CREATE USER 'ews'@'localhost' IDENTIFIED BY 'ews';
GRANT ALL ON ewsdb.* TO 'ews'@'localhost';
FLUSH PRIVILEGES;
exit
```
- Create Tables and Example-Data
```
cd $GOPATH/src/github.com/nlack/ews-qr-app/restapi
mysql -u ews -p < schema.sql
```

### Build Project
`cd $GOPATH/src/github.com/nlack/ews-qr-app/restapi`
`go build api.go`  

### run project
`./api`
