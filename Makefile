APP_NAME=deplk8s
DOCKER_IMG=deplk8s:latest
K8S_NAMESPACE=default

build:
	go build -o $(APP_NAME) main.go

docker-build:
	docker build -t $(DOCKER_IMG) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMG)

k8s-deploy:
	kubectl apply -f deploy/k8s/

k8s-delete:
	kubectl delete -f deploy/k8s/
