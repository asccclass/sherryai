APP?=app
ImageName?=sherry/ai
ContainerName?=sherryAI
PORT?=11001
MKFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURDIR := $(dir $(MKFILE))
GoMode?=on

clean:
	go clean

tidy:
	go mod tidy

build: tidy clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=${GoMode} go build -tags netgo -o ${APP}

buildw:tidy clean
	set GOOS=windows 
	set GOARCH=amd64 
	set CGO_ENABLED=0 
	go build -tags netgo -o ${APP}

docker: build
	docker build -t ${ImageName} .
	rm -f ${APP}
	docker images

log:
	docker logs -f -t --tail 20 ${ContainerName}

stop:
	docker stop ${ContainerName}

rund:docker
	docker run -d --restart=always --name ${ContainerName} \
	-v ${CURDIR}www:/app/www  \
	-v ${CURDIR}envfile:/app/envfile  \
	-v ${CURDIR}tmp:/app/tmp  \
	-p ${PORT}:80 \
	--env-file ${CURDIR}envfile \
	${ImageName}

run: buildw
	./${APP}

re:
	re: stop run

s:
	git push -u origin master