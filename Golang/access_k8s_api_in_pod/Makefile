export tag=1.0.2

hello:
	echo "Go client call k8s api in a pod"

build: hello
	echo "building docker container"
	docker build -t wadexu007/exercise-goclient:${tag} .

push: build
	echo "pushing to my gcr"
	docker push wadexu007/exercise-goclient:${tag}