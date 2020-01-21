
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
* [Docker file Examples](#DockerFileExample)
* [Docker Compose Examples](#DockerComposeExample)
* [Conversion tool Kompose](#ConversionToolKompose)
* [Automation using Codeship](#AutomationUsingCodeship)
* [Maintaining Application Upgrades](#MaintainingApplicationUpgrades)
* [Scaling Application](#ScalingApplication)
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
    curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl

    For example, to download version v1.17.0 on Linux, type:

    curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl

    Make the kubectl binary executable.

    chmod +x ./kubectl

    Move the binary in to your PATH.

    sudo mv ./kubectl /usr/local/bin/kubectl

    Test to ensure the version you installed is up-to-date:

    kubectl version --client
```
## Code

Gorilla/mux package is required to be imported in the application in order to implemet a request router and dispatcher.

import "github.com/gorilla/mux"

Referring to gorilla/mux (https://github.com/gorilla/mux) where it implements a request router and dispatcher for matching incoming requests to their respective handler. 

''

## Docker file Examples

(https://github.com/ContainerSolutions/cd-with-docker-tutorial)

## Docker Compose Examples


## Conversion tool Kompose

## Automation using Codeship

## Maintaining Application Upgrades

## Scaling Application

(https://blog.codeship.com/continuous-integration-and-delivery-with-docker/)

# Authors

* **Kareem Mostafa** - [KareemMostafa](https://github.com/Kareemabdallah)

## Built With

* [Go](https://golang.org/doc/) - The web framework used
* [Python](https://docs.python.org/3/) - Dependency Management
* [Docker](https://docs.docker.com/) - Used to generate RSS Feeds
* [Docker-Compose](https://docs.docker.com/compose/) - The web framework used
* [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) - Dependency Management
* [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - Dependency Management
* [Kompose](https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/) - Used to generate RSS
* [Gorilla/mux](https://www.gorillatoolkit.org/pkg/mux) - The web framework used
* [CodeShip](https://documentation.codeship.com/) - Dependency Management

