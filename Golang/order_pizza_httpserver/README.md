# What's this folder?

## This is an execrise project about order pizzas

* Write by Golang
* Containerized by Docker
* Deployed to Kubernetes [[Manifests](../../Istio/deployment.yaml)] 
<br>

## Features
* Insert pizzas data
* Insert orders data
* Query pizzas
* Query orders
* Query order by ID

<br>

## How to run locally
```
go run main.go
or  
make local
```

## API

### Healthz check
```
curl -X GET 'http://localhost:8080/healthz'
```

### Insert pizzas data
```
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":1,"name":"Pepperoni","price":12}' | jq
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":2,"name":"Capricciosa","price":10}' | jq
curl -X POST 'http://localhost:8080/pizzas' -d '{"id":3,"name":"Margherita","price":15}' | jq
```

### Insert orders data
```
curl -X POST 'http://localhost:8080/orders' -d '{"pizza_id":1,"quantity":3}' | jq
curl -X POST 'http://localhost:8080/orders' -d '{"pizza_id":2,"quantity":2}' | jq
```

### Query data
```
curl -X GET 'http://localhost:8080/pizzas'
curl -X GET 'http://localhost:8080/orders'
curl -X GET 'http://localhost:8080/orders/1' | jq 
curl -X GET 'http://localhost:8080/orders/2' | jq
```

## Logging
[logutils](https://github.com/hashicorp/logutils) from Hashicorp
```
2022/07/18 18:18:19 [INFO] Red Configuration
2022/07/18 18:18:19 [INFO] Start http server
2022/07/18 18:18:24 [ERROR] Can't read pizzas data from csv
2022/07/18 18:18:34 [INFO] Write pizza record to csv
2022/07/18 18:18:39 [INFO] get all pizzas
```