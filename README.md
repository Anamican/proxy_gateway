# Proxy Gateway

## Simple Proxy Gateway in Golang

The purpose of this script to run as a proxy gateway so we could identify the anonymity level of Proxy i.e Elite, Anonymous, Transparent

## Article - How to Detect Proxy Anonymity ? - In progress

## Requirements

* Go 1.12

## Binary Executables

If you dont want to build it yourselves the executables for Mac, Windows and Linux are present in `build` folder

The files are built by running

```shell script
  env GOOS=darwin  GOARCH=amd64 go build -o ./build/proxy_gateway
  env GOOS=linux   GOARCH=amd64 go build -o ./build/proxy_gateway
  env GOOS=windows GOARCH=amd64 go build -o ./build/proxy_gateway.exe
```

## Run the file by

```shell script
    
    go build
    ./proxy_gateway 8080 

```

If you not sure what free port you have in your local system, use the below script  

https://github.com/Anamican/gort to find it automatically

You can go into https://github.com/Anamican/gort/tree/master/cmd/gort folder and run the following:

```shell script    
    go build
``` 

Now you have a `gort` executable and a `proxy_gateway` executable which can piped like this:

```shell script
    gort | xargs -I{} proxy_gateway {}
```

which will return output something like

```shell script
    Proxy Gateway Started on port: 52278
```
