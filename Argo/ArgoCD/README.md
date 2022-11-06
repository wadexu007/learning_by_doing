## What Is Argo CD?

[Argo CD](https://argo-cd.readthedocs.io/en/stable/) is a declarative, GitOps continuous delivery tool for Kubernetes.


### Requirements
* kubectl v1.19.0+
* [kubectx](https://github.com/ahmetb/kubectx) v0.9.4

### Install Argo CD
```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```


### Access The Argo CD API Server
By default, the Argo CD API server is not exposed with an external IP. To access the API server, choose one of the following techniques to expose the Argo CD API server:


**1. Service Type Load Balancer**

You can change Service Type to Load Balancer `kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'`

**2. Gateway** (Recommend)

Follow the [ingress documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/) on how to configure Argo CD with ingress.

* Use [Ingress-nginx](../Ingress-nginx/) and configure `argocd_ingress.yaml`
* Use [Emissary](../Emissary/)
* Use [Ambassador](../Ambassador/)


Now access `https://argocd.wadexu.cloud`

**3. Port Forwarding**

Kubectl port-forwarding can also be used to connect to the API server without exposing the service.

```
kubectl port-forward svc/argocd-server -n argocd 8080:443
The API server can then be accessed using https://localhost:8080
```

### Download Argo CD CLI
```
brew install argocd
```

The initial password for the admin account is auto-generated and stored as clear text in the field password in a secret named argocd-initial-admin-secret in your Argo CD installation namespace. You can simply retrieve this password using kubectl:
```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

Using the username admin and the password from above, login to Argo CD's IP or hostname:

```
argocd login https://argocd.wadexu.cloud
```
or
```
kubectl port-forward svc/argocd-server -n argocd 8080:443
argocd login https://localhost:8080 --username admin --password <repalce_me> 
```

Change the password using the command:

```
argocd account update-password
```


### Register An External Cluster
This step registers a cluster's credentials to Argo CD, and is only necessary when deploying to an external cluster. When deploying internally (to the same cluster that Argo CD is running in), https://kubernetes.default.svc should be used as the application's K8s API server address.
```
# list context
kubectx

argocd cluster add xxx_context
```

### Create An Application From A Git Repository
**Creating Apps Via CLI**
```
kubectl config set-context --current --namespace=argocd

argocd app create my-app --repo https://github.com/wadexu007/learning_by_doing.git --path Kustomize/demo-manifests/services/demo-app/dev --dest-server https://kubernetes.default.svc --dest-namespace demo
```

**Creating Apps Via UI**

It's straightforward, just `New App`, give your app the name `my-app`, use the project `default`, and leave the sync policy as `Manual`.

Connect the https://github.com/wadexu007/learning_by_doing.git repo to Argo CD by setting repository url to the github repo url, leave revision as HEAD, and set the path to `Kustomize/demo-manifests/services/demo-app/dev`

For Destination, set cluster URL to https://kubernetes.default.svc (or in-cluster for cluster name) and namespace to `demo`.

After filling out the information above, click Create

### Sync (Deploy) The Application
**Syncing via CLI**
```
argocd app get my-app
argocd app sync my-app
```

**Syncing via UI**

Just click SYNC in Applications page.

### What's More
Argo CD supports [several different ways](https://argo-cd.readthedocs.io/en/stable/user-guide/application_sources/) in which Kubernetes manifests can be defined:

* Kustomize applications (my above example)
* Helm charts
* A directory of YAML/JSON/Jsonnet manifests, including Jsonnet.
* Any custom config management tool configured as a config management plugin

<br>