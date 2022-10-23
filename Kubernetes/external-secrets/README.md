## What's in this folder?

This folder contains a collection of Kubernetes manifests used to deploy [Kubernetes External Secrets](https://github.com/external-secrets/kubernetes-external-secrets/tree/master/charts/kubernetes-external-secrets) in kubernetes clusters. The manifests are generated using [Kustomize](https://github.com/kubernetes-sigs/kustomize), which can then be deployed on Kubernetes through our CI/CD pipelines or Terraform Kustomize provider.

### Understanding External Secrets Operator
[Kubernetes External Secrets](https://github.com/external-secrets/kubernetes-external-secrets/tree/master/charts/kubernetes-external-secrets) allows you to use external secret management systems (e.g., GCP/AWS Secrets Manager) to securely add secrets in Kubernetes. Read more about the design and motivation for Kubernetes External Secrets on the [GoDaddy Engineering Blog](https://www.godaddy.com/engineering/2019/04/16/kubernetes-external-secrets/).

### Prerequisites
* [Kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) 3.10.0
* Kubernetes External Secrets 8.1.2
* Google Cloud SDK 397.0.0

### Install with Kustomize
* Quick-start
```
kustomize build external-secrets/sre-mgmt-dev/ > ~/secrets_deploy.yaml
kubectl apply -f ~/secrets_deploy.yaml
```
* Through Terraform Kustomize provider, refer to [my example](../../Terraform/kustomize/README.md)
* You can use Helm to install with Terraform as well, [example](../../Terraform/helm/)

### Backend GCP Secret Manager example
#### Enable [Workload Identity](../../Kubernetes/gke-workload-identity/README.md)

#### Add a secret
Copy output encoded json vaule and paste to GCP Secret Manager Console
```
kubectl create secret -n default tls tls-secret-test \
    --save-config --dry-run=client \
    --key ./xxx.key \
    --cert ./xxx.pem \
    -o jsonpath='{.data}'
```

Or use gcloud command (replace `project_id`):
```
kubectl create secret -n default tls tls-secret-test --save-config --dry-run=client --key ./xxx.key --cert ./xxx.pem  -o jsonpath='{.data}' |  gcloud secrets versions add tls-secret --project=${project_id} --data-file=-
```

#### Usage
Once you have kubernetes-external-secrets installed, you can create an external secret with YAML like the usage_example/tls_secrets.yaml

* name - the name the secret will be saved as in kubernetes
* template - only needed if the secret type is different from basic or password
* key - corresponding GCP secret name
* name - property of the json that the secret will be saved as in the kubernetes json 
* property - if your secret in GCP is saved as json, this will grab the secret value under the specified property of the GCP json
* version - if you need to specify version, this is useful, if unset, this will default to latest
* isBinary - necessary if GCP secret is stored as base64, defaults to false (not base64)

```
kubectl apply -f usage_example/tls_secrets.yaml
```

```
% kubectl get secret -n test                     
NAME                  TYPE                                  DATA   AGE
tls-secret            kubernetes.io/tls   
```

for other type of secret like Opaque, once secret generated you just mount to your app container as a volume to use.
```
spec:
  template:
    spec:
      serviceAccountName: xxx
      containers:
      - name: app
        readinessProbe:
          httpGet:
            path: /v1/heartbeat
            port: 8080
        livenessProbe:
          httpGet:
            path: /v1/heartbeat
            port: 8080
        volumeMounts:
        - name: secret
          mountPath: /usr/app/example-secret.json
          subPath: example-secret.json
          readOnly: true
      volumes:
      - name: secret
        secret:
          secretName: example-secret
```

### Cleanup
```
kubectl delete -f usage_example/tls_secrets.yaml

kubectl delete -f ~/secrets_deploy.yaml
```
