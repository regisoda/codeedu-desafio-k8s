# Maratona Full Cycle

## Desafio: Utilizando K8S


##### 1) Servidor Web - Nginx

```bash
./nginx
```

##### 2)  Configuração do MySQL

```bash
./mysql
```

##### 3)  Desafio GO

- Aplicação GO

```bash
./go
```
- Arquivos configuração K8S

```bash
./k8s
```
##### Docker Image

```bash
docker pull regisoda/go-utilizando-k8s
```

##### Running in Docker

```bash
docker run --name go-k8s -p 8080:8080 -d regisoda/go-utilizando-k8s

curl http://localhost:8080
```