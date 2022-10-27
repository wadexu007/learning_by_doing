## cert-manager
[cert-manager](https://cert-manager.io/) is a powerful and extensible X.509 certificate controller for Kubernetes and OpenShift workloads. It will obtain certificates from a variety of Issuers, both popular public Issuers as well as private Issuers, and ensure the certificates are valid and up-to-date, and will attempt to renew certificates at a configured time before expiry.

### Prerequisites
* kubectl v1.19.0
* cert-manager v1.8.2

### Installation
```
kubectl apply -f cert-manager.yaml
```

### Usage
Issue a Let's Encrypt Certificates with a configuration in `certificate-exercise.yaml` file.
```
kubectl apply -f certificate-exercise.yaml
```

### Troubleshooting
https://cert-manager.io/docs/troubleshooting/acme/#overview