# 🚀 deplk8s - Run Go App Locally on Kubernetes with Minikube

---

## 🧰 Prerequisites

Make sure you have the following installed:

- ✅ [Go](https://golang.org/doc/install)
- ✅ [Docker](https://www.docker.com/products/docker-desktop)
- ✅ [kubectl](https://kubernetes.io/docs/tasks/tools/)
- ✅ [Minikube](https://minikube.sigs.k8s.io/docs/)

---

## 🧪 Step-by-Step Local Deployment

#### 1. minikube start
#### 2. eval $(minikube docker-env) 
#### 3. docker build -t deplk8s:latest . 
#### 4. Ensure installed: docker images | grep deplk8s
#### 5. kubectl apply -f k8s/
#### 6. Ensure pods: kubectl get pods
#### 7. Ensure svc: kubectl get svc
#### 8. Run: minikube service deplk8s-service
