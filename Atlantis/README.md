## What's in this folder?

This folder contains atlantis mainfests for deployment in kubernetes.

Reference: https://www.runatlantis.io/docs/deployment.html#kubernetes-kustomize

#### Prerequisites
* Prepare a github user
* Prepare a [Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-token)
* Use any random string generator to create a [Webhook Secret](https://www.runatlantis.io/docs/webhook-secrets.html)

```
echo -n "" > ghUser
echo -n "" > ghToken
echo -n "" > ghWebhookSecret

kubectl create secret generic atlantis --from-file=ghUser --from-file=ghToken --from-file=ghWebhookSecret

```

#### Deployment
```
kustomize build Atlantis/sre-mgmt-dev > deploy.yaml  

kubectl apply -f deploy.yaml  

```

#### Config Ingress Nginx
Refer to Ingress Nginx [deployment](../Ingress-nginx/sre-mgmt-dev)


#### Github webhook config
Reference: https://www.runatlantis.io/docs/configuring-webhooks.html#github-github-enterprise


#### Atlantis.yaml
Refer to [Atlantis.yaml](../atlantis.yaml) in this mono repo


#### Permimssion
Make sure service account of Atlantis running in GKE (in this example) has full permission to your demo GCP project so that it can manipulate resource in this GCP project.

* GKE default service account use node service account.

* (Optional) for GKE Workload Identity: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
  