## What's in this folder?
Ingress-nginx is an Ingress controller for Kubernetes using NGINX as a reverse proxy and load balancer.

This folder contains a collection of Kubernetes manifests used to deploy [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/) in kubernetes clusters. The manifests are generated using [Kustomize](https://github.com/kubernetes-sigs/kustomize), which can then be deployed on Kubernetes through our CI/CD pipelines or Terraform Kustomize provider.

### Understanding Kustomize
https://kustomize.io/

Quick start from [my example](../../Kustomize/demo-manifests/README.md)

### Prerequisites
* [Kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) 3.10.0
* Ingress Nginx 1.3.0

### Deployment
* Quick-start
Update yaml per your request. 

Because enabled https, so need create a secret first. (Optional) Use [external secrets](../Kubernetes/external-secrets/)
```
kubectl create secret -n dmz tls wade-tls-secret \
  --key ./xxx.key \
  --cert ./xxx.pem
```
then
```
kustomize build ingress-nginx-public/sre-mgmt-dev/ > ~/deploy.yaml
kubectl apply -f ~/deploy.yaml
```
* Through Terraform Kustomize provider, refer to [my example](../../Terraform/kustomize/README.md)
* You can use Helm to install with Terraform as well, [example](../../Terraform/helm/)

### Cleanup
```
kubectl delete -f ~/deploy.yaml 
```

### Multipe Ingress Nginx installed in one cluster

**Apply another Ingress Nginx**
```
# deploy second Ingress Nginx for internal gateway usage
kustomize build ingress-nginx-internal/sre-mgmt-dev/ > ~/deploy-internal.yaml
kubectl apply -f ~/deploy-internal.yaml
```
**Automation**

Through [Terraform Kustomize provider](https://registry.terraform.io/providers/kbst/kustomization/latest/docs), refer to [my example](../../Terraform/kustomize/)
