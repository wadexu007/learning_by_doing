## What is kubectl

The Kubernetes command-line tool, [kubectl](https://kubernetes.io/docs/reference/kubectl/), allows you to run commands against Kubernetes clusters. You can use kubectl to deploy applications, inspect and manage cluster resources, and view logs. 

### Installation
Install kubectl on your local machine. Please refer to [here](https://kubernetes.io/docs/tasks/tools/).

Install kubectx as well, kubectx is a utility to manage and switch between gke contexts.

`brew install kubectx` 

### Connection
[**GKE**](../../Terraform/gke/README.md)
```
gcloud container clusters get-credentials sre-dev-gke --region asia-east2 --project sre-cn-dev
```

[**EKS**](../../Terraform/eks/README.md)
```
aws eks update-kubeconfig --region cn-north-1 --name sre-dev 
```

View current context kube config
```
kubectl config view --minify
```

### Useful Commands

#### Context

```
# List contexts
kubectx

gke_xxx-dev-gke
gke_xxx-staging-gke
gke_xxx-prod-gke
  
# to switch context
kubectx gke_xxx-dev-gke
```
 

#### Pods

```
## Get Pod List, -n is specify name space, or --all-namespaces
$ kubectl get pods -n default

## or 
$ kubectl get po -n default

# describe pod in dmz ns
$ kubectl -n dmz describe pod <pod-name>

# describe pod example
$kubectl -n dmz describe pods -l app=core-uc-fe | grep -i port

## Attach to a pod
kubectl exec -it ${podName} bash -n default

# Attach to a pod example
kubectl exec -it core-msg-api-7c889cfb95-tcw5q bash -n app 

# if pod has multiple container, need specific which container via -c
# below example is app container in app namespace
kubectl exec -it core-user-api-77c68d9846-ldd5f bash -n app -c app
```
 

#### Services

```
# get service
kubectl get svc -n default
```
 

#### Deployments

```
## get deployments yaml, dmz is namespace
$ kubectl get deployments -n dmz

# or
$ kubectl get deploy -n dmz

# get deployments core-uc-fe yaml under dmz namespace
$ kubectl get deploy -o yaml core-uc-fe -n dmz

## Query deployed docker image version via grep
kubectl get deploy -o yaml -n dmz core-uc-fe | grep "image:"


## Rolling restart xxx deployment
kubectl -n dmz rollout restart deployments/xxxx

# delete deployment, pods will also being deleted automatcially
kubectl delete -n default deployment <deployment>
```
 

#### Rollback

```
## Rollback to the previous deployed version, example
kubectl rollout undo deployment/core-portal-fe -n dmz

# get revision list
kubectl rollout history deployment/core-portal-fe -n dmz

# describe the specific revision
kubectl rollout history deployment/core-portal-fe --revision=11 -n dmz

# rollback to specific version
kubectl rollout undo deployment/core-portal-fe --to-revision=9 -n dmz
```


#### Logs

```
## Get Logs from a pod

kubectl logs --tail 1000 -n dmz core-portal-fe-765b4567c7-kf9lv
```
or
```
kubectl logs -f --tail 100  deployment/core-portal-fe  -n dmz
```
 

 

#### List all image name

```
# list all image name
kubectl get pods --all-namespaces -o jsonpath="{..image}" |\
tr -s '[[:space:]]' '\n' |\
sort |\
uniq -c
```
 
 

#### Start a curl client

```
# for re-use this client, remove --rm, otherwise it will be terminated when you exit this pod
$ kubectl run curl-client --image=radial/busyboxplus:curl -i --tty --rm 

# you can use curl command in this pod
```

#### Start a busy box for diagnostics

```
$ kubectl run -i --tty busybox --image=busybox -- sh

$ kubectl attach busybox -c busybox -i -t

# you can execute telnet/nslookup and other commands in this pod
```
Start busybox with google sdk so that you can execute gsutil/gcloud command
```
kubectl run -i --tty busybox --image=google/cloud-sdk:latest -- sh
```

#### Kube Proxy

```
# Example, for your local call core-uc-api service which running in Kubernetes

kubectl -n app port-forward service/core-uc-api 8010:8080

# in your local, call http://127.0.0.1:8010/xxxx
```


<br>