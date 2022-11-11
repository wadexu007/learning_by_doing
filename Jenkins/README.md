## What Is Jenkins

[Jenkins](https://www.jenkins.io/) is an open source automation server which enables developers around the world to reliably build, test, and deploy their software.

## What's in this folder?

* An [installation guide](./install_on_k8s/) for setting up a Jenkins cluster on Kubernetes
* How to [setup Jenkins build agents on Kubernetes Pods](./k8s_pod_as_build_agent/)
  * A [quick example](./k8s_pod_as_build_agent/demo-app-java/) for java app CI (Continuous Integration) using Jenkinsfile as a declarative pipeline
  * A [example](./k8s_pod_as_build_agent/demo-app-go/) for CICD for a golang app (CI + Kustomize + CD to one cluster)
  * This [example](./k8s_pod_as_build_agent/demo-app-go/) also show how to deploy to multiple clusters in different regions with user interaction via Jenkins [Pipelie: Input Step](https://www.jenkins.io/doc/pipeline/steps/pipeline-input-step/) function.
* Build Docker Image In Kubernetes Using Kaniko - [Example](./kaniko-demo/)


<br>