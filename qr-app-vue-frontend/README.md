# qr-app-vue-frontend

> A Vue.js project

Front End to use the qr-backend

## Features
- [x] Login as participant
- [x] Logout as participant
- [x] Login as instructor
- [x] Logout as instructor
- [x] Show participant qr-code
- [x] Show instructor course list
- [x] Show instructor register participant to course window
- [x] Check permission in routing
- [x] Add new Course


## Requirements

- `Node 8.1+`
- Optional run with Docker

## Configuration

Need to change URL to backend server adress!

Edit file `config/prod.env.js`

Replace `URL` in `API_URL: '"URL"'` with you own configuration.

## Run with Docker

Goto vue project folder

`cd $VUE_PROJECT`

Build docker container

`docker build . -t vue-qr-app`

Run docker container

`docker run vue-qr-app:latest`

visit: http://DOCKER_CONTAINER_IP:1337

#### Get docker DOCKER_CONTAINER_IP with the following steps:

Get CONTAINER_ID with:

`docker container ps`

than run:

`docker inspect CONTAINER_ID | grep '"IPAddress"' | head -n 1`

Example output:

`"IPAddress": "172.17.0.2",`

Example address:

`http://172.17.0.2:1337`

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:1337
npm run dev

# build for production with minification, need webserver to! file:// doesn't work!
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```
