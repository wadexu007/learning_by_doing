## Intro

This is a demo app with gin web framework and gorm

## Framework Integration

- [x] HTTP web framework:https://github.com/gin-gonic/gin
- [x] Authentication: gin-jwt middleware: https://github.com/appleboy/gin-jwt
- [x] ORM: https://github.com/go-gorm/gorm


## Dependency Injection
It uses Google wire to manage the dependency injection, Google wire is code generation tool that automates connecting components using dependency injection.

Please read [this article][]  to understand the value of Dependency Ingestion, and why choose Google wire. 

[this article]: https://blog.golang.org/wire

### How to use Google wire?
There is a wire.go in the router package, please add the initialization methods into wire go and then navigate to the router folder in ther terminal, and run the command as below to generate the methods in the wire_gen.go file.
```
# remove comments in wire.go
# delete wire_gen.go then re-create
go run github.com/google/wire/cmd/wire

# add below comments back to wire.go
//go:build wireinject
// +build wireinject

```

## Test in your local
```
# update conf/app_dev.yaml

# if not export env, use app.yaml
export RUN_ENV=dev

make local

```

## API Example

### Common API
#### Health check
```
curl --location --request GET 'http://localhost:8080/health'

pong!

```

#### Generate Token

```
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"admin",
    "password":"xxxx"
}'

```

#### Refresh Token

```
curl --location --request GET 'http://localhost:8080/auth/refresh_token' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY5NDIwMTYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2Njk0MDIxNn0.YKEXW-8oEfPyz_bComawj5VWPc-dhdEUY7NA8_Uwjj4'

```


#### Create

```
curl --location --request POST 'http://localhost:8080/v1/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY5NDIwMTYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2Njk0MDIxNn0.YKEXW-8oEfPyz_bComawj5VWPc-dhdEUY7NA8_Uwjj4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"wadexu",
    "email":"wade.xu@demo.com"
}'

"b686b42c-6ccb-47d2-a6de-bc3f588ec2cf"
```

#### Search by name
```
curl --location --request GET 'http://localhost:8080/v1/account/search?name=wadexu' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY5NDIwMTYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2Njk0MDIxNn0.YKEXW-8oEfPyz_bComawj5VWPc-dhdEUY7NA8_Uwjj4' | jq

[
  {
    "id": "7524ec4e-8fbb-4bdb-b9db-bd69ea980064",
    "createdAt": "2022-10-28T14:58:42.961+08:00",
    "updatedAt": "2022-10-28T14:58:42.961+08:00",
    "deletedAt": null,
    "name": "wadexu",
    "email": "wade.xu@demo.com",
    "createdBy": "system",
    "updateBy": ""
  }
]

```

#### Update

```
curl --location --request PUT 'http://localhost:8080/v1/account/7524ec4e-8fbb-4bdb-b9db-bd69ea980064' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY5NDIwMTYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2Njk0MDIxNn0.YKEXW-8oEfPyz_bComawj5VWPc-dhdEUY7NA8_Uwjj4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"wadexu",
    "email": "wade.xu@new.com"
}'

```

#### Delete
```
curl --location --request DELETE 'http://localhost:8080/v1/account/a2227041-c99b-4833-a5b1-35e6cee5e084' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY5NDIwMTYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2Njk0MDIxNn0.YKEXW-8oEfPyz_bComawj5VWPc-dhdEUY7NA8_Uwjj4' 

```
