# What's this repo?
Helm-go-client utilize Helm SDK to manage the application deployment on Kubernetes, this is Go application to provide the [features](#features).

## 
* Write by Golang
* Containerized by Docker
* Manage Kubernetes [[YAML manifests](../../Kustomize/demo-manifests/services/helm-go-client/)] by Kustomize
* Continuous Integration by [[Github Actions](../../GitOps/github_actions/cicd-helm-go.yaml)]
* Continuous Deployment by ArgoCD
<br>

## Features
* Switch K8s cluster and run helm actions on the specific cluster.
* List Helm releases
* Get Helm releases
* Install Helm charts
* Delete Helm releases
* Update Helm charts (only resource cpu/mem)
* Update Helm charts (any value)

## Kube Config
Before run this application, need make sure kube config is ready, config path specify in `./conf/config.json`.
<br>
Here is kube config example for two clusters:
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: xxx
    server: https://xx.xx.10.2
  name: sre-dev-gke
- cluster:
    certificate-authority-data: xxx
    server: https://xx.xx.10.1
  name: sre-spark-test-gke
contexts:
- context:
    cluster: sre-dev-gke
    user: sre-dev-gke-token-user
  name: sre-dev-gke
- context:
    cluster: sre-spark-test-gke
    user: sre-spark-test-gke-token-user
  name: sre-spark-test-gke
kind: Config
users:
- name: sre-dev-gke-token-user
  user:
    token: xxx
- name: sre-spark-test-gke-token-user
  user:
    token: xxx
```

## Add a new cluster
### Generate a service account user token
Token is for above kube config user's token.
```
# connect to new cluster locally
CONTEXT=$(kubectl config current-context)

# This service account uses the ClusterAdmin role, more restrictive roles can by applied.
kubectl apply --context $CONTEXT -f ./manifests/service-account.yaml

TOKEN=$(kubectl get secret --context $CONTEXT \
   $(kubectl get serviceaccount helm-service-account \
       --context $CONTEXT \
       -n ops \
       -o jsonpath='{.secrets[0].name}') \
   -n ops \
   -o jsonpath='{.data.token}' | base64 --decode)
```

## How to run locally
```
make local
```

## Continuous Deployment
[[Github Actions](../.github/workflows/cicd-helm-go.yaml)]
* Trigger by push/pull_request and files changes in current folder
* Build and Push to GCR automatically
* Kustomize edit image tag and push changes
* Trigger ArgoCD auto sync deployment

## API

### Healthz check
```
curl 'http://localhost:8080/healthz'
```

### List Helm release in a namespace
List all Helm release in `sre-dev-gke` cluster `default` namespace
```
% curl 'http://localhost:8080/list?clusterContext=sre-dev-gke&namespace=default' | jq 
[
  {
    "name": "ingress-dmz",
    "namespace": "default",
    "revision": 1,
    "status": "deployed",
    "chart_path": "ingress",
    "app_version": "1.16.0"
  },
  {
    "name": "loki",
    "namespace": "default",
    "revision": 1,
    "status": "deployed",
    "chart_path": "loki-stack",
    "app_version": "v2.4.2"
  }
]
```

### List Helm release in all namespace
List all Helm release in `sre-spark-test-gke` cluster All namespace
```
% curl 'http://localhost:8080/list?clusterContext=sre-spark-test-gke' | jq
[
  {
    "name": "jupyterhub-dev",
    "namespace": "default",
    "revision": 10,
    "status": "deployed",
    "chart_path": "jupyterhub",
    "app_version": "1.5.0"
  },
  {
    "name": "spark-dev",
    "namespace": "default",
    "revision": 15,
    "status": "deployed",
    "chart_path": "spark",
    "app_version": "3.2.1"
  }
]
```

### Get specific Helm release
Get `sealed-secrets` Helm release in `sre-dev-gke` cluster `kube-system` namespace
```
% curl 'http://localhost:8080/get?clusterContext=sre-dev-gke&namespace=kube-system&name=sealed-secrets' | jq 
{
  "name": "sealed-secrets",
  "namespace": "kube-system",
  "revision": 2,
  "status": "deployed",
  "chart_path": "sealed-secrets",
  "app_version": "v0.18.1"
}
```

### Install Helm chart
Install Helm chart `bitnami/nginx` with release name `nginx-demo` in `sre-dev-gke` cluster `default` namespace
```
% curl -X POST 'http://localhost:8080/install' -d '{"cluster":"sre-dev-gke", "chartPath":"./helm-repos/nginx-13.1.5.tgz","namespace":"default","name":"nginx-demo"}' | jq
{
  "name": "nginx-demo",
  "namespace": "default",
  "revision": 1,
  "status": "deployed",
  "chart_path": "nginx",
  "app_version": "1.23.1"
}
```

### Update Helm Chart only resouce value
Update Helm release `nginx-demo` with specific resource `request cpu/memory` and upgrade chart version to `13.1.6` in `sre-dev-gke` cluster `default` namespace
```
% curl -X PUT 'http://localhost:8080/updateOnlyResource' -d '{"input": {"cluster":"sre-dev-gke", "chartPath":"./helm-repos/nginx-13.1.6.tgz","namespace":"default","name":"nginx-demo"}, "values": {"request_cpu":"150m","request_memory":"128Mi"}}' | jq
```
after update check result via `kubectl get deploy nginx-demo -oyaml`
```
        resources:
          requests:
            cpu: 150m
            memory: 128Mi
```

### Update Helm Chart any value
e.g. update `replicaCount`
```
% curl -X PUT 'http://localhost:8080/updateHelmAnyValue' -d '{"cluster":"sre-dev-gke", "chartPath":"./helm-repos/nginx-13.1.6.tgz","namespace":"default","name":"nginx-demo","values":{"replicaCount":3}}' | jq
```

### Delete Helm release
Delete `nginx-demo` Helm release in `sre-dev-gke` cluster `default` namespace
```
% curl -X DELETE 'http://localhost:8080/delete' -d '{"cluster":"sre-dev-gke","namespace":"default","name":"nginx-demo"}' | jq
```

<br>
