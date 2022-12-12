export TAG=1.1.3
export REPO=wadexu007/instrumented-app

hello:
	echo "Welcome to Bookstore app with instrumentation."

local:
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${REPO}:${TAG} .

push: build
	echo "pushing to my gcr"
	docker push ${REPO}:${TAG}
