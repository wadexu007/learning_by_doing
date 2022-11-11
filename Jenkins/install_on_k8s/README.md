## Prerequisites
1. A running kubernetes cluster (My example use GKE 1.22)

## Setup Jenkins On Kubernetes Cluster
For setting up a Jenkins cluster on Kubernetes, we will do the following.

* Create a namespace
* Create a service account with Kubernetes admin permissions.
* Create a Persistent Volume using PVC on GKE for persistent Jenkins data on Pod restarts.
* Create a deployment YAML and deploy it.
* Create a service YAML and deploy it.
* Access the Jenkins application.

### Deployment
Step 1: Create a Namespace for Jenkins.
```
kubectl apply -f namespace.yaml
```

Step 2: Create k8s service account and RBAC permission
```
kubectl apply -f serviceAccount.yaml
```

Step 3: Create a Persistent Volume using PVC on GKE
* Create a storage class
* Provision a Persistent volume using the storage class.

```
kubectl apply -f volume.yaml
```
Check result:
```
kubectl get pvc -n jenkins

NAME              STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
jenkins-storage   Bound    pvc-27efe7b9-c963-4366-b100-a3b01bb25666   20Gi       RWO            jenkins-sc     23s
```

Step 4: Create Deployment
Make persistent volume to hold Jenkins data path /var/jenkins_home is very important for production use cases, otherwise data will lost once pods restarts.
```
kubectl apply -f deployment.yaml
```
Check the deployment status.
```
kubectl get deploy -n jenkins

NAME      READY   UP-TO-DATE   AVAILABLE   AGE
jenkins   1/1     1            1           89s
```

Step 5: Create Service
```
kubectl apply -f service.yaml
```


### Access Jenkins
* Option 1 (recommend): Use a gateway

 1. **[Emissary Ingress](../../Emissary/)**
 2. **[Ingress-nginx](../../Ingress-nginx/ingress-nginx-public/sre-mgmt-dev/)**

* Option 2: Use Kube Proxy
```
kubectl -n jenkins port-forward service/jenkins-service 8010:8080
```
Then in your local browser, you will be able to access the Jenkins dashboard via http://127.0.0.1:8010

<br>
Jenkins will ask for the initial Admin password when you access the dashboard for the first time.

```
kubectl get pods -n jenkins

kubectl logs jenkins-998474795-7n6ls -n jenkins
```
```
*************************************************************

Jenkins initial setup is required. An admin user has been created and a password generated.
Please use the following password to proceed to installation:

xxxxxxxxxxxxxx

This may also be found at: /var/jenkins_home/secrets/initialAdminPassword
```

Once you enter the password you can proceed to install the suggested plugin and create an admin user. All these steps are self-explanatory from the Jenkins dashboard.


## Reference:
* https://github.com/scriptcamp/kubernetes-jenkins/blob/main/deployment.yaml
* https://devopscube.com/setup-jenkins-on-kubernetes-cluster/
* https://devopscube.com/persistent-volume-google-kubernetes-engine/
* https://www.jenkins.io/doc/book/installing/kubernetes/


## HA
* Jenkins active/passive setup - only **enterprise** Jenkins comes with a supported plugin to have this setup.

* This doc's way, Jenkins running on Kubernetes where if the Jenkins pod goes down, another pod will come up with the same data. 

When you host Jenkins on Kubernetes for production workloads, you need to consider setting up a highly available persistent volume to avoid data loss during pod or node deletion. Another way using a nfs server (e.g. GCP [filestore](https://cloud.google.com/filestore)) for data stored.

Create a filestore first, e.g. server IP is `10.127.2.110`

`pv and pvc yaml`
```
---
# pv
apiVersion: v1
kind: PersistentVolume
metadata:
  name: gcp-filestore-jenkins
spec:
  capacity:
    storage: 100Gi
  accessModes:
  - ReadWriteMany
  nfs:
    path: /vol1/jenkins_home
    server: 10.127.2.110
---
# pvc
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: gcp-filestore-jenkins-claim
  namespace: jenkins
spec:
  # Specify "" as the storageClassName so it matches the PersistentVolume's StorageClass.
  # A nil storageClassName value uses the default StorageClass. For details, see
  # https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class-1
  accessModes:
  - ReadWriteMany
  storageClassName: ""
  volumeName: gcp-filestore-jenkins
  resources:
    requests:
      storage: 100Gi
```

`example part of deployment.yaml`
```
    spec:
      terminationGracePeriodSeconds: 10
      serviceAccountName: jenkins-admin
      securityContext:
            fsGroup: 1000 
            runAsUser: 1000
      initContainers:
          - name: init-folder-permssion
            image: busybox:latest
            command: ["sh","-c","mkdir -p /var/jenkins_home && chown 1000:1000 /var/jenkins_home"]
            volumeMounts:
              - name: gcp-filestore-jenkins
                mountPath: /var/jenkins_home
      containers:
        - name: jenkins
          image: jenkins/jenkins:lts
          ports:
            - containerPort: 8080
            - containerPort: 50000
          volumeMounts:
          - mountPath: /var/jenkins_home
            name: gcp-filestore-jenkins
      volumes:  
      - name: gcp-filestore-jenkins
        persistentVolumeClaim:
          claimName: gcp-filestore-jenkins-claim
          readOnly: false
```

<br>