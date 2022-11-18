export TAG=1.0.6
export REPO=wadexu007/exercise-pizza

hello:
	echo "Welcome to pizza order system."

local:
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${REPO}:${TAG} .

push: build
	echo "pushing to my gcr"
	docker push ${REPO}:${TAG}
