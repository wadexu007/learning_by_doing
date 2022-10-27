# Demo how to access k8s API in a Pod via go client

[Offical docs](https://kubernetes.io/docs/tasks/run-application/access-api-from-pod/#accessing-the-api-from-within-a-pod)
<br>
[Go client library](https://github.com/kubernetes/client-go/)


## Init project
```
go mod init main.go
go mod tidy  
```

## Deploy
This will push to my perosnal repository, replace to yours.
```
# edit tag in Makefile
make push

# edit image tag in deployment.yaml
kubectl apply -f deployment.yaml
```

## Arguments in deployment.yaml
For look up dmz and app namespaces
```
    args:
    - dmz
    - app
```

## Test
```
2022/07/23 14:18:22 There are 26 pods in namespace dmz of the cluster
2022/07/23 14:18:22 Name: ingress-nginx-controller-696b489f67-n7qxp-controller, Version: k8s.gcr.io/ingress-nginx/controller:v1.2.0@sha256:d8196e3bc1e72547c5dec66d6556c0ff92a23f6d0919b206be170bc90d5f9185
......
......
2022/07/23 14:18:22 There are 47 pods in namespace app of the cluster
2022/07/23 14:18:22 Name: ambassador-679495d67c-ctt9k-ambassador, Version: docker.io/datawire/ambassador:1.12.1
2022/07/23 14:18:22 Name: ambassador-679495d67c-glkc4-ambassador, Version: docker.io/datawire/ambassador:1.12.1
2022/07/23 14:18:22 Name: ambassador-679495d67c-ktvzx-ambassador, Version: docker.io/datawire/ambassador:1.12.1
2022/07/23 14:18:22 Name: lr-nginx-859d8f47bb-ld4wm-my-nginx, Version: nginx
```