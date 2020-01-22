
# Containerized microservice challenge

This challenge implements the following: 

* It implements an application written in golang for exposing a JSON document 'db.json'.
* It implements another application written in python for reversing the text acquired from the first Application.
* It dockerizes the two applications in Dockerfiles.
* Dockerfile could be written into docker-compose files for more automation.
* Docker-compose files could be converted into Kubernetes Resources using Kompose conversion tool.
    https://kompose.io/
    https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/
    
* After conversion applications could run into a Minikube Kubernetes installed locally.
* it automates the two applications with bash scripts using go language.

---

* [Prerequisites](#Prerequisites)
* [Code](#Code) (Go, Python)
* [Kompose](#Kompose)
* [Automation](#Automation)
* [Maintaining Applications Upgrades](#MaintainingApplicationUpgrades)
* [Scaling Applications](#ScalingApplications)
* [Authors](#Authors)
* [Built With](#BuiltWith)


## Prerequisities

1. Golang:
```
go get -u github.com/gorilla/mux

```

2. Docker:
```
https://docs.docker.com/install/
```
3.Minikube:
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube
  sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
```
4. Kubectl:
```
    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl
   
   
    chmod +x ./kubectl

    sudo mv ./kubectl /usr/local/bin/kubectl
    
    kubectl version --client
```
## Code

Gorilla/mux package is required to be imported in the application in order to implemet a request router and dispatcher.

import "github.com/gorilla/mux"

Referring to gorilla/mux (https://github.com/gorilla/mux) where it implements a request router and dispatcher for matching incoming requests to their respective handler. 

## Kompose

## Automation using Codeship

## Maintaining Application Upgrades

## Scaling Application

(https://blog.codeship.com/continuous-integration-and-delivery-with-docker/)

# Authors

* **Kareem Mostafa**

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

