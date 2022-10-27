## What's in this folder?

This folder contains some best practices for hardening your containers, including:

* Limit Capabilities and Privilege Escalation
* Secure Containers Using Linux Security Modules
* Disable Inter-Container Communication
* Limit the container to just the necessary resources


### Limit Capabilities and Privilege Escalation
The principle of least privilege also applies when running containers. Linux allows about 100 capabilities that can be added or dropped when running a container. The entire capabilities list can be found [here](https://man7.org/linux/man-pages/man7/capabilities.7.html).

By default, Docker runs with a subset of these privileges. When running a docker container, you can drop privileges (`--cap-drop`) or add privileges (`--cap-add`). The most secure setup is to drop all privileges and then add privileges back on only as required.

Let's try this out by first running our Docker image:

```
docker run -it ubuntu bash
```

Exit out of this container by typing `exit` then run it again with all of the permissions removed:

```
docker run --rm -it --cap-drop all ubuntu bash
```

Once the docker container is running, try creating a new user named wadexu with this command:
`adduser wadexu`

You should see a failure message.
```
Adding user `wadexu' ...
Adding new group `wadexu' (1000) ...
groupadd: failure while writing changes to /etc/gshadow
adduser: `/sbin/groupadd -g 1000 wadexu' returned error code 10. Exiting.
```

Now exit out of this container by typing `exit` and add back the capabilities needed to create a user and change the permissions for a file:

```
docker run --rm -it --cap-drop all --cap-add AUDIT_WRITE --cap-add CHOWN --cap-add DAC_OVERRIDE --cap-add FOWNER ubuntu bash
```

After the container starts, create the `wadexu` user and change the ownership of a file with the following commands:

Add a new user: `adduser wadexu`

Change ownership of a file: 
```
touch server.js

chown wadexu:wadexu server.js
```

Check ownership with: `ls -la` 

These commands should show that the new owner of the server.js file is wadexu:

```
-rw-r--r-- 1 wadexu wadexu    0 Oct 27 13:06 server.js
```


Also, remember to be wary of using the `--privileged` flag, which will add all the Linux Kernel capabilities to the container. "Privileged" containers are often used in the development phase to speed up development, but should be avoided in production.

It is also important to prevent escalation of privileges while running your container. This is done by setting the flag to --security-opt=no-new-privileges. This will prevent the processes inside the container from gaining new privileges during execution, such as running sudo or su. This would look something like this:

```
docker run --security-opt no-new-privileges -it ubuntu bash
```

### Secure Containers Using Linux Security Modules
* AppArmor

[AppArmor](https://docs.docker.com/engine/security/apparmor/) is a Linux Security Module that you can configure to restrict capabilities of processes running in the container. AppArmor can allow or deny a subject's access to an object. For example, AppArmor can allow a program read-only access to `/etc/hosts` but have no access to `/etc/passwd`. AppArmor has its own [default profile](https://github.com/moby/moby/blob/master/profiles/apparmor/template.go) that you can customize. To run AppArmor, use the `--security-opt` flag in your run command:

```
docker run --rm -it --security-opt apparmor=docker-default --name node_apparmor ubuntu bash
```

* seccomp
[seccomp](https://docs.docker.com/engine/security/seccomp/) is another option for restricting the actions available within a container. Where AppArmor focuses on subject-object access, seccomp focuses on syscalls available, such as adding kernel keyrings `add_key`, setting the date/time `clock_adjtime`, or mounting `mount`. seccomp has a [default profile](https://github.com/moby/moby/blob/master/profiles/seccomp/default.json) that you can customize. seccomp is specified when running the container by defining the profile as follows:

```
docker run --rm -it --security-opt seccomp=seccomp.json --name node_seccomp ubuntu bash
```

AppArmor, seccomp, and dropping privileges can all be used in conjunction to reduce the damage of a compromised container.


### Disable Inter-Container Communication

In order to set up and run an application, usually many containers will need to be created. However, not all of those containers need to communicate with each other. For example, the backend database may not need to communicate directly with the load balancer. By default, when setting up the docker environment, all containers can talk to each other, unless you specifically set the environment otherwise with the `--icc=false` flag. This is done when setting up the docker daemon with the following command: `dockerd --icc=false`. This will prevent direct communication between all containers unless communication is specifically stated.

NOTE: The Docker daemon cannot be restarted in this sandbox, but you can test this feature on your host machine.

To create a connection between containers, you will create a user-defined [bridge network](https://docs.docker.com/network/bridge/#differences-between-user-defined-bridges-and-the-default-bridge) with a command such as:

```
docker network create my-network
```
You would then connect a running container to that network with the command:
```
docker network connect my-network my-container-name
```


### Limit the container to just the necessary resources

Applications may unexpectedly consume large amounts of compute resources, either from an attacker or from a bug in the application. Limitations on containers can help keep these problems under control. Take a look at the following command and we will examine each option:
```
docker run -it --cpus=".5" --memory 128m --restart=on-failure:5 ubuntu bash
```

* --memory flag: defines the maximum amount of memory the container can use. In our example, if the container allocates more than 128Mb, then the request will be denied
* --cpus flag: specifies how much of the available CPU resources a container can use. In our example, if you have 1 CPU, then this container will use at most 50% of the CPU every second.
* --restart flag: this flag defines the restart policy of the container. The default value is no. In our example, we state the restart policy to on-failure and we define the maximum number of restarts as 5. This will prevent the container from an ongoing loop of crashing and restarting.

There must be enough resources for the container to function, but not too many resources to allow a Denial of Service attack. The limitations will need to be balanced based on the needs of the container.


## Conclusion
These are some of the many recommendations for container hardening. You can find more best practices from many organizations including [OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Docker_Security_Cheat_Sheet.html), [Snyk](https://snyk.io/blog/10-docker-image-security-best-practices/), and [Docker Security](https://docs.docker.com/engine/security/) . It will be up to your team to decide which hardening techniques are best for your application.

<br>