export TAG=1.0.7
export PROJECT=main.go
# docker hub container registry
export DCR=wadexu007

hello:
	echo "This is a demo"

local: hello
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${DCR}/${PROJECT}:${TAG} .

push: build
	echo "pushing to my gcr"
	docker push ${DCR}/${PROJECT}:${TAG}