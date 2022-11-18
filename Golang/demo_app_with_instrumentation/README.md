## Bookstore REST API using Gin and Gorm
Read this [article](https://blog.logrocket.com/rest-api-golang-gin-gorm/) on understand how to build a sample golang app using Gin and Gorm.

### Local Run
Comment out Tracer part in `main.go`
```
go mod tidy

go run main.go
```

### Test API
```
GET    /books                    
GET    /books/:id               
POST   /books                    
PATCH  /books/:id                
DELETE /books/:id                
```

* Insert books
```
curl -X POST 'http://localhost:8080/books' -d '{"title":"Cloud Native","author":"wadexu"}' | jq
curl -X POST 'http://localhost:8080/books' -d '{"title":"Linux","author":"wadexu"}' | jq
```

* Get books
```
curl -X GET http://localhost:8080/books | jq
```

## Send Trace data to Grafana Tempo
Refer to [here](../../Grafana/Grafana%20Tempo/)