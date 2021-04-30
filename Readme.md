# Users Api using Golang and mySQl

- This is an API that is part of a bigger microservice

## To build the container

```bash 
docker build -t main .
```

## To run the container

```bash 
docker run -p 8080:8080 -p 9200:9200 users-api:latest
```