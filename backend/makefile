VERSION := 0.0.1 
APP_NAME := xeroaccounting

build:
	docker build \
	-t ${APP_NAME}:${VERSION} \
	--build-arg VERSION=${VERSION} \
	. 

run:
	docker compose up --force-recreate --build
	
test:
	docker run -t \
	--volume $$(pwd):/go/src/accounting \
	--rm  \
	$$(docker build --quiet --file test.Dockerfile . ) 

lint: 
	docker run --rm \
	--volume $$(pwd):/src \
	--volume ~/.cache:/root./.cache \
	$$(docker build --quiet --file lint.Dockerfile .) 

