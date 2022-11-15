export TAG=1.0.0
export PROJECT=helm-go-client
export GCR=asia.gcr.io/sre-dev
export TAG2?=-none

hello:
	echo "This is Go client call helm sdk"

local: hello
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${GCR}/${PROJECT}:${TAG}${TAG2} .

push: build
	echo "pushing to my gcr"
	docker push ${GCR}/${PROJECT}:${TAG}${TAG2}