## Overview

[GitHub Actions](https://docs.github.com/en/actions) is a continuous integration and continuous delivery (CI/CD) platform that allows you to automate your build, test, and deployment pipeline. You can create workflows that build and test every pull request to your repository, or deploy merged pull requests to production.


## Quick Start
https://docs.github.com/en/actions/quickstart


## My Example

### Workflow for polyrepo
This is a [workflow pipeline](./.github_example/workflows/cicd-helm-go.yaml) for my golang app [helm-go-client](../Golang/helm-go-client/) with [deployment mainfests](../Kustomize/demo-manifests/services/helm-go-client/).

**Workflows breakdown**
1. Checkout repo
2. Setup go env
3. Login gcr
4. Docker build and push to gcr
5. Setup Kustomize
6. Checkout k8s manifests repo
7. Update Kubernetes resources with kustomize edit
9. Git push changes
10. Trigger [ArgoCD](../ArgoCD/)


### Manually running a workflow
Reference: [workflow_dispatch](https://docs.github.com/en/actions/managing-workflow-runs/manually-running-a-workflow)

[my example](./.github_example/workflows/workflow_dispatch.yaml)


### Workflow for monorepo
[my example](./.github_monorepo/workflows/cicd.yaml)

**Workflows breakdown**
1. Checkout repo
2. Set env
3. Get changed files
4. Set matrix for build
5. Setup go
6. Login gcr
7. Docker build and push to gcr
9. Checkout k8s manifests repo
10. Update Kubernetes resources with kustomize edit
11. Set up Cloud SDK
12. Set up GKE credentials
13. Deploy to the GKE cluster
14. Commit image version to track

<br>
