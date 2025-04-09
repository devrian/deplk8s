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

### 1️⃣ Start Minikube

1. eval $(minikube docker-env) 
2. docker build -t deplk8s:latest . 
3. Ensure installed: docker images | grep deplk8s
4. kubectl apply -f k8s/
5. Ensure pods: kubectl get pods
6. Ensure svc: kubectl get svc
7. Run: minikube service deplk8s-service
