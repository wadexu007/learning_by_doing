## What's in this file?

This file contains useful Helm commands

### Install
Install bitnami nginx as example
```
helm repo add bitnami https://charts.bitnami.com/bitnami

helm repo list

helm search repo nginx

helm install mywebserver bitnami/nginx

helm list

# install to a specific namespace
kubectl create ns demo

helm install mywebserver bitnami/nginx -n demo

helm list -n demo

# install with a specific nginx version
helm install mywebserver bitnami/nginx --version 13.2.8

```

#### change default values then install

```
helm inspect values bitnami/nginx > values.yaml

vim values.yaml

helm install mywebserver bitnami/nginx -f values.yaml

NAME: mywebserver
LAST DEPLOYED: Tue Oct  4 10:23:16 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: nginx
CHART VERSION: 13.2.9
APP VERSION: 1.23.1
```

### Pull
download nginx chart from bitnami repository and (optionally) unpack it in local directory
```
helm pull bitnami/nginx --untar

# install from local
helm install mywebserver2 ./nginx 
```

### Clean Up
```
helm list

NAME        	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART       	APP VERSION 
mywebserver2	default  	1       	2022-10-04 10:40:35.587977 +0800 CST	deployed	nginx-13.2.9	1.23.1
mywebserver 	default  	1       	2022-10-04 10:43:53.863981 +0800 CST	deployed	nginx-13.2.9	1.23.1
```

```
helm uninstall mywebserver

helm repo remove bitnami
```