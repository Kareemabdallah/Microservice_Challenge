[![Build Status](https://travis-ci.org/Kareemabdallah/Microservice_Challenge.svg?branch=master)](https://travis-ci.org/Kareemabdallah/Microservice_Challenge)

# Containerized microservice

This repossitory implements the following: 

* Implements an application written in golang exposing a JSON document 'static/db.json' when visiting HTTP client.
* Implements another go application which reverses the message text acquired from the first Application.
* Dockerizes the two applications.
* Docker-compose file is added when using docker-compose.
* Docker-compose file can be converted to a k8s deployment for minikube implementations.
* Build script for building the docker images and run it.
* After conversion could run into a Minikube Kubernetes installed locally.
* Scaling out.
* Maintains the regular applications upgrades.

---

* [Prerequisites](#Prerequisites)
* [Dependencies](#Dependencies)
* [Built With](#BuiltWith)


## Prerequisities

1. Go:
```
https://golang.org/doc/install
```

2. Docker:
```
https://docs.docker.com/install/
```

3. Docker Compose:
```
https://docs.docker.com/compose/install/
```
4. Minikube:
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube
  sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
```
5. Kubectl:
```
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl
   
   
    chmod +x ./kubectl

    sudo mv ./kubectl /usr/local/bin/kubectl
    
    kubectl version --client
```
6. Kompose
```
    https://kompose.io/
    
    https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/
   ```

## Dependencies

Gorilla/mux package is required to be imported inside the application in order to implemet a request router and dispatcher by adding:

```
import "github.com/gorilla/mux"
```
It must be downloaded at first using the command:

```
go get -u github.com/gorilla/mux
```

After cloning the Repository, we can run the the two applications using docker-compose:

```
docker-compose up
```

## Built With

* [Go](https://golang.org/doc/)
* [Python](https://docs.python.org/3/) 
* [Docker](https://docs.docker.com/) 
* [Docker-Compose](https://docs.docker.com/compose/) 
* [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) 
* [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) 
* [Kompose](https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/) 
* [Gorilla/mux](https://www.gorillatoolkit.org/pkg/mux)
* [CodeShip](https://documentation.codeship.com/) 
