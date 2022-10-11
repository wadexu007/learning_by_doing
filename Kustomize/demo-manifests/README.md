## What's in this folder?

This folder contains a collection of Kubernetes manifests used to deploy an example [demo app](../../Docker/demo/) workloads in kubernetes clusters. The manifests are generated using [Kustomize](https://github.com/kubernetes-sigs/kustomize), which can then be deployed on Kubernetes through our CI/CD pipelines

### Understanding Kustomize
https://kustomize.io/


### Install Kustomize
https://kubectl.docs.kubernetes.io/installation/kustomize/

### Usage
* Use kustomize command

```
cd Kustomize/demo-manifests

kustomize build services/demo-app/dev/  > deploy.yaml 

```

* Use [kubectl](https://kubernetes.io/docs/reference/kubectl/) command 

Kustomize already natively built into kubectl since 1.14

```
kubectl kustomize services/demo-app/dev/ > deploy.yaml
```

### Test
Create kubernetes resource from yaml file
```
kubectl apply -f deploy.yaml
```

### Cleanup
```
kubectl delete -f deploy.yaml 
```
