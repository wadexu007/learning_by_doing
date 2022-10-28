export TAG=1.0.0
export PROJECT=helm-go-client
# docker hub container registry
export DCR=wadexu007
export TAG2?=-none

hello:
	echo "This is Go client call helm sdk"

local: hello
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${DCR}/${PROJECT}:${TAG}${TAG2} .

push: build
	echo "pushing to my gcr"
	docker push ${DCR}/${PROJECT}:${TAG}${TAG2}