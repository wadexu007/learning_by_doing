## What Is Atlantis?
[Atlantis](https://www.runatlantis.io/) is an application for automating Terraform via pull requests. It is deployed as a standalone application into your infrastructure. No third-party has access to your credentials.

This folder contains atlantis mainfests for deployment in kubernetes.

Reference: https://www.runatlantis.io/docs/deployment.html#kubernetes-kustomize

## Installation
#### Prerequisites
* A running kubernetes cluster
* Prepare a github user
* Prepare a [Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-token)
* Use any random string generator to create a [Webhook Secret](https://www.runatlantis.io/docs/webhook-secrets.html)

```
echo -n "xxx" > ghUser
echo -n "xxxx" > ghToken
echo -n "xxxxx" > ghWebhookSecret

kubectl create ns atlantis
kubectl create secret generic atlantis --from-file=ghUser --from-file=ghToken --from-file=ghWebhookSecret -n atlantis
```

**Github App integration refer to my [medium](https://medium.com/@wadexu007/c74bce4c7fde?source=friends_link&sk=21b5112a96c1b244fd7e9e47ffd1c00e).**

#### Permimssion
Make sure service account of Atlantis running in GKE (in this example) has full permission to your demo GCP project so that it can manipulate resource in this GCP project.

* GKE default service account use node service account.

* For GKE Workload Identity: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity

Workload Identity allows workloads in your GKE cluster to impersonate GSA (Google Service Accounts) using KSA (Kubernetes Service Accounts) configured during deployment.

Create GSA for use with Workload Identity
```
SA_NAME="atlantis"
GKE_PROJECT_ID="adept-presence-396401"
SA_EMAIL="$SA_NAME@${GKE_PROJECT_ID}.iam.gserviceaccount.com"

gcloud iam service-accounts create $SA_NAME --display-name $SA_NAME

gcloud projects add-iam-policy-binding $GKE_PROJECT_ID \
--member serviceAccount:$SA_EMAIL --role "roles/owner"
```

Link KSA to GSA
```
#namespace in kubernetes
NS="atlantis"

gcloud iam service-accounts add-iam-policy-binding $SA_EMAIL \
  --role "roles/iam.workloadIdentityUser" \
  --member "serviceAccount:$GKE_PROJECT_ID.svc.id.goog[$NS/atlantis]"

```

Check binding result
```
gcloud iam service-accounts get-iam-policy $SA_EMAIL

#result
bindings:
- members:
  - serviceAccount:adept-presence-396401.svc.id.goog[atlantis/atlantis]
  role: roles/iam.workloadIdentityUser
etag: BwYIHSdPl-Y=
version: 1
```

#### Deployment
```
kustomize build sre-mgmt-dev | kubectl apply -f -
```

#### Config Ingress Nginx
Refer to Ingress Nginx [deployment](../Ingress-nginx/ingress-nginx-public/sre-mgmt-dev/)


#### Github webhook config
Reference: https://www.runatlantis.io/docs/configuring-webhooks.html#github-github-enterprise


#### Atlantis.yaml
Refer to [Atlantis.yaml](../atlantis.yaml) in this monorepo

## Workflow Test
* Step 1: Open a Pull Request
* Step 2: Atlantis automatically run `terraform plan` and comments back on PR
* Step 3: Someone reviews and approves PR
* Step 4: Comment `atlantis apply`
* Step 5: Atlantis run `terraform apply` and comments back on PR about result
* Step 6: PR merged automatically.

![alt text.](../Images/atlantis_auto_plan_terraform_PR.jpg "This is test result image.")

More details
[Getting Started with Atlantis on GKE for Terraform Automation with GCP Workload Identity](https://medium.com/@wadexu007/8b5ca1884d05?source=friends_link&sk=58f2ad8bf8d7c720a34628b100ce2ed7)

<br>
