## What's in this folder?
[Emissary Ingress](https://www.getambassador.io/docs/emissary/) API gateway from Ambassador Labs is an open source Kubernetes-native API gateway built on the Envoy Proxy.

This folder contains a collection of Kubernetes manifests used to deploy [Emissary Ingress](https://www.getambassador.io/docs/emissary/) in kubernetes clusters. The manifests are generated using [Kustomize](https://github.com/kubernetes-sigs/kustomize), which can then be deployed on Kubernetes through our CI/CD pipelines or Terraform Kustomize provider.

### Understanding Kustomize
https://kustomize.io/

Quick start from [my example](../../Kustomize/demo-manifests/README.md)

### Prerequisites
* [Kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) 3.10.0
* Emissary-Ingress 3.2.0

### Deployment
* Quick-start

Update yaml in dev folder per your request
```
kustomize build quick-start/emissary-ingress/sre-mgmt-dev/ > ~/emissary_deploy.yaml

kubectl apply -f ~/emissary_deploy.yaml
```

* Through Terraform Kustomize provider, refer to [my example](../../Terraform/kustomize/README.md)


### Cleanup
```
kubectl delete -f ~/emissary_deploy.yaml 
```

### More example
[Multiple emissary installed in one cluster](./multiple-emissary-example/)

From emissary-ingress 2.1, it removed CRDs from its charts, refer to official installation docs, apply CRDs is first step now. `kubectl apply -f https://app.getambassador.io/yaml/emissary/3.2.0/emissary-crds.yaml`

That's why I make a kustomize init folder for install Emissary CRDs first.

**Notes: Install multiple Emissary in one cluster must set ambassador_id and replace ClusterRoleBinding name.** 
Check details code in emissary-ingress-private/public to see.

* emissary-ingress-init CRDs will be installed.
* emissary-ingress-public An emissary-ingress with allow list = all (face to internet).
* emissary-ingress-private Another emissary-ingress with an allow list (restrict connection) installed in same cluster.


**Test in local**
```
# apply CRDs first
kustomize build emissary-ingress-init/sre-mgmt-dev > ~/init.yaml
kubectl apply -f ~/init.yaml

# deploy first public Emissary, this allow list = all, face to internet
kustomize build emissary-ingress-public/sre-mgmt-dev > ~/emissary_deploy1.yaml
kubectl apply -f ~/emissary_deploy1.yaml

# deploy second private Emissary with a restrict allow list to access
kustomize build emissary-ingress-private/sre-mgmt-dev > ~/emissary_deploy2.yaml
kubectl apply -f ~/emissary_deploy2.yaml
```
**Automation**

Through [Terraform Kustomize provider](https://registry.terraform.io/providers/kbst/kustomization/latest/docs), refer to [my example](../../Terraform/kustomize/)
