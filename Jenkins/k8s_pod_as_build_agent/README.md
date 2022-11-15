## Setup Jenkins Build Agents on Kubernetes Pods

### Prerequisites
1. A running Jenkins master in a Kubernetes cluster (my example use GKE 1.22) refer to [Installation Guide](../install_on_k8s/)
2. A Service Account with k8s admin in a target cluster for deployment
3. Firewall allowed Jenkins outbound to Docker Hub
4. Firewall to target cluster from Jenkins should be allowed.

### Installed Plugin
* Kubernetes Plugin
* Google Kubernetes Engine Plugin (For deployment to GKE cluster)

### Configuration
1. go to `Manage Jenkins` –> `Manage Nodes and Clouds`
2. Click `Configure Clouds`
3. Add a new Cloud select `Kubernetes`
4. Click `Kubernetes Cloud Details`
5. Enter `jenkins` namespace in `Kubernetes Namespace` field
6. Click `Test Connection` --> result show `Connected to Kubernetes v1.22.12-gke.2300`
7. Click `Save`
8. Enter `http://jenkins-service.jenkins.svc.cluster.local:8080` in `Jenkins URL` field
9. Enter `jenkins-agent:50000` in `Jenkins tunnel` field
10. Click `Add Pod Template` then `Pod Template Details`
11. Input `Name`=`jenkins-agent`, `Namespace`=`jenkins`, `Labels`=`kubeagent`
12. (Optional)  If you don’t add a container template, the Jenkins Kubernetes plugin will use the default JNLP image from the Docker hub to spin up the agents. But if your Jenkins can't access Docker Hub. Then build your own `jnlp` image and override the default with the same name. Click `Add Container` to add Container Template, `Name`=`jnlp`, `Docker Image`=`your_registry/jenkins/inbound-agent:4.11-1-jdk11`, **Ensure that you remove the sleep and 9999999 default argument from the container template**.

### Manage Credentials
* Add `Usernames with password` for docker hub account/pwd
* Add `Google Service Account from private key`
Replace Jenkinsfile - environment variable to yours.

## Test a freestyle project
Go to Jenkins home –> New Item and create a freestyle project.

In the job description, add the label `kubeagent` for `Restrict where this project can be run`. It is the label we assigned to the pod template. This way, Jenkins knows which pod template to use for the agent container.

Add a shell build step with an echo command
```
echo "testing"
```

Now, save the job configuration and click “Build Now”

In a short while, you will see a successful build.


## Declarative-pipeline
Use Jenkinsfile With Pod Template, the Pod template is defined inside the kubernetes { } block.
* **Reference**:
 * A quick start for java app CI (demo-app-java/Jenkinsfile)
 * CICD for a golang app  (demo-app-go/Jenkinsfile)
 
The demo-app-go example is more complex with below scenarios:
* CI + Kustomize build k8s YAMLs + Deploy to one cluster
* Deploy to multiple clusters in different regions with user interaction via Jenkins [Pipeline: Input Step](https://www.jenkins.io/doc/pipeline/steps/pipeline-input-step/) function.


**Now you can create a Pipeline or Multibranch Pipeline job in Jenkins to test.**
1. Repository URL = `https://github.com/wadexu007/learning_by_doing`
2. Script Path, e.g. `Jenkins/k8s_pod_as_build_agent/demo-app-java/Jenkinsfile`


## Conclusion
At this point, you have successfully created a CI/CD pipeline between Jenkins, GitHub, Docker Hub, Kustomize and one or more Kubernetes clusters running on Google Kubernetes Engine. You can now extend and enhance it with [multiple branches](https://www.jenkins.io/doc/tutorials/build-a-multibranch-pipeline-project/) and [notifications](https://www.jenkins.io/doc/pipeline/tour/post/).

## Best Practice
* **Docker login**

Don't use below way to login docker registry. 
```
docker.withRegistry("$DOCKER_HUB_REGISTRY", "$DOCKER_HUB_CREDENTIAL") {}
```
Refer to below insecure msg, check my example `demo-app-go` use `--password-stdin`
```
[Pipeline] withDockerRegistry
Executing sh script inside container docker of pod wade-app-05fxg-zsg91
Executing command: "docker" "login" "-u" "wadexu007" "-p" ******** "https://index.docker.io/v1" 
exit
WARNING! Using --password via the CLI is insecure. Use --password-stdin.
```

* **Mount Maven m2 cache directory**

When you build the java apps, it downloads dependencies added in the pom.xml from the remote maven repository the first time, and it creates a local .m2 cache directory where the dependant packages are cached.

The .m2 cache is not possible in Docker agent-based builds as it gets destroyed after the build.

We can create a persistent volume for the maven cache and attach it to the agent pod via container template to solve this issue.

Create pvc and mount it maven container `_common/build-pod.yaml`
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: maven-repo-storage
  namespace: jenkins
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Gi
```

* Set resource requests and limits on each container in your Pod

* Try use [Jenkins Shared Libraries](https://www.jenkins.io/doc/book/pipeline/shared-libraries/) for common part of Pipelines

* **Run docker in docker**
Using [kaniko](../kaniko-demo/) instead of using docker.sock method which is less secure as it has complete privileges over the docker daemon.

## FAQ

* **Issue 1**: Came along this issue `Jenkins: 403 No valid crumb was included in the request` when I changed Jenkins to be accessible via reverse proxy. There is an option in the "Configure Global Security" that "Enable proxy compatibility" will address it.


* **Issue 2**: I met below issue when not do `Configuration` - `Step 9` about Jenkins tunnel
```
2022-11-09 10:44:06.124+0000 [id=96]	WARNING	o.c.j.p.k.KubernetesLauncher#launch: Error in provisioning; agent=KubernetesSlave name: kube-agent-21kt8, template=PodTemplate{id='13c16c61-de41-4e35-bd1c-4294190b8419', name='kube-agent', namespace='jenkins', label='kubeagent', containers=[ContainerTemplate{name='jnlp', image='jenkins/inbound-agent:4.3-4', workingDir='/home/jenkins/agent', command='', args='', resourceRequestCpu='', resourceRequestMemory='', resourceRequestEphemeralStorage='', resourceLimitCpu='', resourceLimitMemory='', resourceLimitEphemeralStorage='', livenessProbe=ContainerLivenessProbe{execArgs='', timeoutSeconds=0, initialDelaySeconds=0, failureThreshold=0, periodSeconds=0, successThreshold=0}}]}
java.lang.IllegalStateException: Node was deleted, computer is null
	at org.csanchez.jenkins.plugins.kubernetes.KubernetesLauncher.launch(KubernetesLauncher.java:191)
	at hudson.slaves.SlaveComputer.lambda$_connect$0(SlaveComputer.java:298)
	at jenkins.util.ContextResettingExecutorService$2.call(ContextResettingExecutorService.java:48)
	at jenkins.security.ImpersonatingExecutorService$2.call(ImpersonatingExecutorService.java:82)
	at java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)
	at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)
	at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)
	at java.base/java.lang.Thread.run(Thread.java:829)
```

* **Issue 3**: Build agent always pending. 
if you add container template for `jnlp` in Console, ensure that you remove the sleep and 9999999 default argument from the container template.
```
2022-11-09 11:48:54.903+0000 [id=559]	INFO	o.c.j.p.k.KubernetesLauncher#launch: Waiting for agent to connect (30/1,000): jenkins-agent-h3kdl
2022-11-09 11:49:25.094+0000 [id=559]	INFO	o.c.j.p.k.KubernetesLauncher#launch: Waiting for agent to connect (60/1,000): jenkins-agent-h3kdl
2022-11-09 11:49:55.291+0000 [id=559]	INFO	o.c.j.p.k.KubernetesLauncher#launch: Waiting for agent to connect (90/1,000): jenkins-agent-h3kdl
```

* **Issue 4**: Job failed with below error,  need add `git config --global --add safe.directory`
```
fatal: detected dubious ownership in repository at '/home/jenkins/agent/workspace/test-wade_master'
To add an exception for this directory, call:

	git config --global --add safe.directory /home/jenkins/agent/workspace/test-wade_master
```

* **Issue 5**: Job failed due to `Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?`
Your docker container doesn't contains Docker daemon(engine), to resolve it, mount `/var/run/docker.sock` in k8s host.
Running docker in docker using docker.sock method is less secure as it has complete privileges over the docker daemon.
Refer to [kaniko](../kaniko-demo/) deployment is a better way.


* **Issue 6**: Make sure firewall rule is opened. Allow Jenkins access to Target k8s cluster on port 443.
```
Logs: Unable to connect to the server: dial tcp 34.92.xx.xxx:443: i/o timeout
```

* **Issue 7**: Make sure grant jenkins user/group to deployment.yaml `chown 1000:1000 deployment.yaml`
```
Also:   hudson.remoting.Channel$CallSiteStackTrace: Remote call to JNLP4-connect connection from 172.20.72.150/172.20.72.150:44864
......
......
java.nio.file.AccessDeniedException: /home/jenkins/agent/workspace/xxx/xxx/demo-app-go/deployment.yaml
```

## Useful links
https://devopscube.com/jenkins-build-agents-kubernetes/
https://www.jenkins.io/doc/book/pipeline/syntax/#declarative-pipeline
https://www.jenkins.io/doc/pipeline/steps/pipeline-input-step/


## Known issue
[Google Kubernetes Engine Plugin](https://github.com/jenkinsci/google-kubernetes-engine-plugin/blob/develop/docs/Home.md#usage) - verifyDeployments(boolean) does not work for non-default namespace https://github.com/jenkinsci/google-kubernetes-engine-plugin/issues/297

<br>