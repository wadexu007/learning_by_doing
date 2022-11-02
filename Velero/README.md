## What Is Velero

[Velero](https://velero.io/docs/v1.9/index.html) gives you tools to back up and restore your Kubernetes cluster resources and persistent volumes. You can run Velero with a cloud provider or on-premises.

## Installation
### Environment Info
* GKE 1.22
* Kustomize 3.10.0

### CLI Tool
Reference: [Basic Install](https://velero.io/docs/v1.9/basic-install/)

I'm using MacOS with below command
```
brew install velero
```

### Cloud Provider
I'm using GCP as example.
* [Manually](https://github.com/vmware-tanzu/velero-plugin-for-gcp#setup)
* [Terraform](./gcp_resource/sre-mgmt-dev/) - Recommend
```
cd gcp_resource/sre-mgmt-dev
terraform init
terraform plan 
terraform apply
```

### Deployment
* Deploy with [kustomize](../Kustomize/demo-manifests/). Modify kustomize_install/base folder these files: `backupstoragelocation.yaml` and `serviceaccount.yaml` per your requeset then execute:
```
kustomize build kustomize_install/velero/sre-mgmt-dev > ~/deploy.yaml
kubectl apply -f ~/deploy.yaml
```

* (Optional) Deploy with [Helm](https://github.com/vmware-tanzu/helm-charts/tree/main/charts/velero) manually, try [Terraform Helm Provider](../Terraform/helm/) if you want.


## Backup
### [Schedule API](https://velero.io/docs/v1.9/api-types/schedule/)
Velero have a Schedule resource that uses cron format to continuously backup the cluster at the specified time. My Velero deployment kustomize yaml includes a schedule that backups the cluster hourly. To create additional schedules; either create additional YAML files or use the velero schedule create command.
```
# Create a backup every 6 hours.
$ velero create schedule NAME --schedule="0 */6 * * *"

# Create a backup every 6 hours with the @every notation.
$ velero create schedule NAME --schedule="@every 6h"

# Create a daily backup of the web namespace.
$ velero create schedule NAME --schedule="@every 24h" --include-namespaces web

# Create a weekly backup, each living for 90 days (2160 hours).
$ velero create schedule NAME --schedule="@every 168h" --ttl 2160h0m0s

# Schedule help
$ velero create schedule --help
```

### [Backup API](https://velero.io/docs/v1.9/api-types/backup/)
Having a Schedule will create the Backup resource at the scheduled time. Backups can be created using YAML or velero backup create command.
```
# Create a backup containing all resources.
$ velero backup create backup1

# Create a backup including only the nginx namespace.
$ velero backup create nginx-backup --include-namespaces nginx

# Create a backup excluding the velero and default namespaces.
$ velero backup create backup2 --exclude-namespaces velero,default

# Create a backup based on a schedule named daily-backup.
$ velero backup create --from-schedule daily-backup

# View the YAML for a backup that doesn't snapshot volumes
$ velero backup create backup3 --snapshot-volumes=false

# Backup help
$ velero backup create --help
```

## Disaster Recovery (DR) 
Restoring specific resources
```
$ velero backup get  # pick the backup to restore from

# selector is lables in cluster resource
$ velero restore create RESTORE_NAME --from-backup BACKUP_NAME --selector key1=value1,key2=value2
```

Restoring namespaces
```
$ velero backup get  # pick the backup to restore from

$ velero restore create RESTORE_NAME --from-backup BACKUP_NAME --include-namespaces namespace1,namespace2
```

Complete cluster restore

```
$ velero backup get  # pick the backup to restore from

$ velero restore create RESTORE_NAME --from-backup BACKUP_NAME
```

<br>
