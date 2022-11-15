
## Introduction
Previously, we build docker in docker using mounting `docker.sock` to host way, the Docker build containers run in privileged mode. It is a big security concern and it is kind of an open door to malicious attacks. You check out this folder [demo-app-go](../k8s_pod_as_build_agent/) to understand more.

[kaniko](https://github.com/GoogleContainerTools/kaniko) is an open-source container image-building tool created by Google. It does not require privileged access to the host for building container images.

## Build Docker Image In Kubernetes Using Kaniko
### Prerequisites
1. Jenkins running on Kubernetes cluster
2. Docker Hub account ready
3. Firewall allowed Jenkins outbound to Docker Hub

### Create Dockerhub Kubernetes Secret
We have to create a kubernetes secret of type docker-registry for the kaniko pod to authenticate the Docker hub registry and push the image in `jenkins` namespace.
```
kubectl -n jenkins create secret docker-registry dockercred \
    --docker-server=https://index.docker.io/v1/ \
    --docker-username=<dockerhub-username> \
    --docker-password=<dockerhub-password>
```

### Define Kaniko pod in yaml
Refer to [build-pod.yaml](build-pod.yaml)
```
    - name: kaniko
      image: gcr.io/kaniko-project/executor:v1.9.0-debug # include shell
      imagePullPolicy: IfNotPresent
      command:
        - /busybox/cat
      tty: true
      resources:
        limits:
          cpu: 500m
          memory: 1024Mi
      volumeMounts:
      - name: kaniko-secret
        mountPath: /kaniko/.docker
  volumes:
  - name: kaniko-secret
    secret:
        secretName: dockercred
        items:
        - key: .dockerconfigjson
          path: config.json
```

### Jeknis Pipeline with Kaniko
```
    stage('Kaniko Build and Push Docker Image') {
      steps {
        script {
          dir(dir_path) {
            container('kaniko') {
              sh """
                /kaniko/executor --context `pwd` --destination wadexu007/$PROJECT_IMAGE_WITH_TAG
              """
            }
          }
        }
      }
    }
```

* --context: This is the location of the Dockerfile. 
* --destination: Here, you need to replace dockerhub-username `wadexu007` with your docker hub username for kaniko to push the image to the dockerhub registry.


## Known Issue
https://github.com/GoogleContainerTools/kaniko/issues/1586#issuecomment-895349111