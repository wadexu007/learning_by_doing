## What Is Kubernetes?
Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications.

## General Practices
* [kubectl](./kubectl/README.md) List all useful commands.
* [lb-service](./lb-service/) Example spec for loadbalancer type service in EKS

## Security Practices
* [external-secrets](./external-secrets/) Demo how to use external secret management systems (GCP Secrets Manager) to securely add secrets in Kubernetes.
* [cert-manager](./cert-manager/) Demo how to use cert-manager to issue a certificate from let's Encrypt.
* [gke-workload-identity](./gke-workload-identity/) This is the recommended way for your workloads running on GKE to access Google Cloud services in a secure and manageable way.
* [kubescape](./kubescape/) Scanning your k8s manifest files to search for common vulnerabilities that may be lurking in your code.

<br>
