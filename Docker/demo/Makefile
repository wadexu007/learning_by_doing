export TAG=1.0.0
export REPOSITORY=wadexu007/demo

hello:
	echo "Welcome to order pizza system."

local:
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${REPOSITORY}:${TAG} .

push: build
	echo "pushing to my gcr"
	docker push ${REPOSITORY}:${TAG}
