## Intro

This is a gorm [has many](https://gorm.io/docs/has_many.html) association example.

This project using same framework as [webframework-gin](../webframework-gin/), read it first.

### Test in your local
```
# update conf/app_dev.yaml

# if not export env, use app.yaml
export RUN_ENV=dev

make local

```

### API Example

#### Common API
##### Health check
```
curl --location --request GET 'http://localhost:8080/health'

pong!

```

##### Generate Token

```
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"admin",
    "password":"xxxx"
}'

```

##### Refresh Token

```
curl --location --request GET 'http://localhost:8080/auth/refresh_token' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjcwMjk5NTksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2NzAyODE1OX0.E8t4Cq-71VMeWIrl-HAzRKkZuLtoK4ZAG_3c5pXHWec'

```


##### Create

```
curl --location --request POST 'http://localhost:8080/v1/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjcwMjk5NTksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2NzAyODE1OX0.E8t4Cq-71VMeWIrl-HAzRKkZuLtoK4ZAG_3c5pXHWec' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"wadexu",
    "email":"wade.xu@demo.com",
    "creditCards": [
      {
        "number": "6666666"
      },
      {
        "number": "8888888"
      }
    ]
}'

"b686b42c-6ccb-47d2-a6de-bc3f588ec2cf"
```

##### Search by name
```
curl --location --request GET 'http://localhost:8080/v1/account/search?name=wadexu' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjcwMjk5NTksImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2NzAyODE1OX0.E8t4Cq-71VMeWIrl-HAzRKkZuLtoK4ZAG_3c5pXHWec' | jq

[
  {
    "id": "538f05d6-9508-4cc6-9918-265657f3c6a5",
    "createdAt": "2022-10-29T15:24:07.339+08:00",
    "updatedAt": "2022-10-29T15:24:07.339+08:00",
    "deletedAt": null,
    "name": "wadexu",
    "email": "wade.xu@demo.com",
    "createdBy": "system",
    "updateBy": "",
    "creditCards": [
      {
        "ID": 1,
        "CreatedAt": "2022-10-29T15:24:07.911+08:00",
        "UpdatedAt": "2022-10-29T15:24:07.911+08:00",
        "DeletedAt": null,
        "number": "6666666",
        "UserID": "538f05d6-9508-4cc6-9918-265657f3c6a5"
      },
      {
        "ID": 2,
        "CreatedAt": "2022-10-29T15:24:07.911+08:00",
        "UpdatedAt": "2022-10-29T15:24:07.911+08:00",
        "DeletedAt": null,
        "number": "8888888",
        "UserID": "538f05d6-9508-4cc6-9918-265657f3c6a5"
      }
    ]
  }
]
```

### Test result
```
mysql> select * from user;
+--------------------------------------+-------------------------+-------------------------+-------------------------+--------+-------------------+------------+------------+
| id                                   | created_at              | updated_at              | deleted_at              | name   | email             | created_by | updated_by |
+--------------------------------------+-------------------------+-------------------------+-------------------------+--------+-------------------+------------+------------+
| 538f05d6-9508-4cc6-9918-265657f3c6a5 | 2022-10-29 15:24:07.339 | 2022-10-29 15:38:55.063 | 2022-10-29 15:38:57.174 | wadexu | wade.xu@demo.com  | system     | system     |
+--------------------------------------+-------------------------+-------------------------+-------------------------+--------+-------------------+------------+------------+
1 rows in set (0.17 sec)
```
```
mysql> select * from credit_card;
+----+-------------------------+-------------------------+------------+----------+--------------------------------------+
| id | created_at              | updated_at              | deleted_at | number   | user_id                              |
+----+-------------------------+-------------------------+------------+----------+--------------------------------------+
|  1 | 2022-10-29 15:24:07.911 | 2022-10-29 15:24:07.911 | NULL       | 6666666  | 538f05d6-9508-4cc6-9918-265657f3c6a5 |
|  2 | 2022-10-29 15:24:07.911 | 2022-10-29 15:24:07.911 | NULL       | 8888888  | 538f05d6-9508-4cc6-9918-265657f3c6a5 |
+----+-------------------------+-------------------------+------------+----------+--------------------------------------+
2 rows in set (0.17 sec)

```
<br>

