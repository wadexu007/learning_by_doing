export TAG=1.0.6
export PROJECT=exercise-pizza
export CR=wadexu007
export TAG2?=-none

hello:
	echo "Welcome to pizza order system."

local:
	echo "run locally"
	go run main.go

build: hello
	echo "building docker container"
	docker build -t ${CR}/${PROJECT}:${TAG}${TAG2} .

push: build
	echo "pushing to my gcr"
	docker push ${CR}/${PROJECT}:${TAG}${TAG2}
