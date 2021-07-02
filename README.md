# docker-first-app

- source code available [here](https://github.com/HuanLiu-hotstar/docker-first-app)

## create a repo for demo

- with go.mod format

## create a server  or app

- we can write a server with golang
- and listen on port:`9090`

## create a Dockerfile in server repo

- Dockerfile used to generate an image
- server will run in container based image

```Docker

FROM golang:1.16 # inherit the base images

WORKDIR /app ## location the server locate

COPY go.mod . # copy go.mod to workdir 
# COPY go.sum .

RUN go mod download ## in golang we should download some dependencies to compile the server 

COPY . . ## copy all files to workdir, maybe some configure files

## compile source code & generate the executable file
## this will generate an executable file named  main ,details see Makefile
## you can also use other shell comand to compile server
RUN GOOS=linux GOARCH=amd64  make   

## start the server 
ENTRYPOINT ["/app/main"] ## start server command 

```

## build image

```shell
# 1. build image for docker engine 

docker build -t golang-test:v1 . 
# -t: give a name for image 
# .: directory,contains Dockerfile 

# 2. we can search the image in docker engine
docker images 

## 3. run an instance of the image 
docker run -d -p 5000:9090 golang-test:v1

# -d: means run the instance in backgroud
# -p: map localhost port to instance port in docker engine
# last one parameters is image name

## 4.verify the service 
curl localhost:5000/  

```

## congratulation

- you finished the hello world for Docker deploy with your server