## What's in this folder?
Ambassador API gateway from Ambassador Labs is an open source Kubernetes-native API gateway built on the Envoy Proxy.

Now it called [Emissary Ingress](https://www.getambassador.io/docs/emissary/), if you want to use latest Emissary Ingress, pls refer to [this example](../Emissary/). To still use old Ambassador, refer to below.

This folder contains a collection of Kubernetes manifests used to deploy Ambassador in kubernetes clusters. The manifests are generated using [Kustomize](https://github.com/kubernetes-sigs/kustomize), which can then be deployed on Kubernetes through our CI/CD pipelines or Terraform Kustomize provider.

### Understanding Kustomize
https://kustomize.io/

Quick start from [my example](../../Kustomize/demo-manifests/README.md)

### Prerequisites
* [Kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) 3.10.0
* Ambassador 1.12.1

### Deployment
* Quick-start

Update yaml in dev folder per your request
```
kustomize build gateway-public/dev/ > ~/deploy.yaml

kubectl apply -f ~/deploy.yaml
```

* Through Terraform Kustomize provider, refer to [my example](../../Terraform/kustomize/README.md)


* Cleanup
```
kubectl delete -f ~/deploy.yaml
```

### Multiple Ambassador installed in one cluster

* gateway-internal
* gateway-public

**Notes: Install multiple Ambassador in one cluster must set ambassador_id and replace ClusterRoleBinding name.** 
Check details code in gateway-internal and gateway-public folder to see.

**Test in local**
```
# deploy public Ambassador, this allow list = all, face to internet
kustomize build gateway-public/dev/ > ~/public_gateway.yaml
kubectl apply -f ~/public_gateway.yaml

# deploy second internal Ambassador which is for internal LB
kustomize build gateway-internal/dev/ > ~/internal_gateway.yaml
kubectl apply -f ~/internal_gateway.yaml
```
**Automation**

Through [Terraform Kustomize provider](https://registry.terraform.io/providers/kbst/kustomization/latest/docs), refer to [my example](../../Terraform/kustomize/)

**Cleanup**
```
kubectl delete -f ~/public_gateway.yaml

kubectl delete -f ~/internal_gateway.yaml
```

<br>