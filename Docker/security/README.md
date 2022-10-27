### Docker Security - Limit Size of Containers

#### Prefer Minimal Base Images
Each docker image starts with a base image. Often the features in popular base images will contain features that are necessary for your goals. For example, popular images for running a postgres database can be based on a variety of different base images. Understanding which images to use requires an investigation of the features provided by each image. The features can range from a full blown operating system with unnecessary tools and libraries for your organization’s purpose.

#### Use "Distroless" Container Images for Production
Often it’s efficient to have features and tools present in docker images that allow developers to debug, monitor and modify the docker-based systems they are developing. But as the systems are promoted to production these tools can increase the surface-area of the system.

"Distroless" Container Images remove these tools and are available for many environments requiring nodejs, java, python, etc. You can learn more about ["Distroless" Container Images here](https://github.com/GoogleContainerTools/distroless).


#### Use Multistage Builds

Often docker images build as well as run the application. It’s possible using multistage builds to allow for one base image for building the application and another image for running the application. For example, say you wish to compile and run a java application. You could use the following Dockerfile:

```
# Example 1
FROM openjdk:11`
COPY ./HelloWorld.java .
RUN javac ./HelloWorld.java
CMD java -cp . HelloWorld
```
But to run a java program we don’t need to have the java development kit installed. We only need the java runtime environment. We could update our dockerfile as such:

```
# Example 2
FROM openjdk:11
COPY ./HelloWorld.java .
RUN javac ./HelloWorld.java

FROM openjdk:11-jre
COPY --from=0 ./HelloWorld.class .
CMD java -cp . HelloWorld
```
The resulting docker image in example 2 would be about half of the image size of the docker image in example 1.

Go example refer to [here](../demo/)

#### Reduce Image Size using Optimization Tools
Using tools such as [Docker Slim](https://dockersl.im/) or [Dive](https://github.com/wagoodman/dive) can reduce the size of your image by discovering unneeded tools or libraries.

<br>