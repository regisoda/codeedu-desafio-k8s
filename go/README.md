## Build 

```bash
go build main.go
```

## Run Test

```bash
go test
```

## Docker Image

```bash
docker pull regisoda/go-utilizando-k8s
```

## Running in Docker

```bash
docker run --name go-k8s -p 8080:8080 -d regisoda/go-utilizando-k8s

curl http://localhost:8080
```