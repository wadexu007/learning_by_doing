
## What is this folder
This folder demos how containerizing a Go applications with Docker Multistage Builds and then push image to docker hub registry.

### Use Multistage Builds
[Dockerfile](./Dockerfile)

### Build
```
make build
```
or 
```
docker build -t wadexu007/demo:1.0.0
```

### Docker hub
* Register an account in  https://hub.docker.com

* Account Settings --> Security --> New Access Tokens

* Generate a new one. Then docker login, replace `wadexu007` with your account name.
```
docker login -u wadexu007
```
after paste access token: `Login Succeeded`

* (Optional) You can also edit `~/.docker/config.json` directly:
```
% cat ~/.docker/config.json
{
    "auths": {
        "asia.gcr.io": {
            "auth": "xxx"
        },
        "https://index.docker.io/v1/": {
            "auth": "xxx"
        },
        "xxx-cn-north-1.jcr.service.jdcloud.com": {
            "auth": "xxx"
        }
    }
}% 
```
The second one is docker hub `index.docker.io/v1`



* Create a `demo` repository in docker hub console.
* Push image
```
make push
```
or
```
docker push wadexu007/demo:1.0.0
```


#### Reference
[My Blog](https://www.cnblogs.com/wade-xu/p/16708050.html)

<br>