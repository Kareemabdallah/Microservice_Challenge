
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

go get -u github.com/gorilla/mux

2. 
Sudo apt-get update -y  && \
     apt-get install -y python3.7-dev && \
     apt-get install -y python-pip && \
     pip install -y flask

Sudo pip3 install -r requirements.txt

virtualenv flask


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

* **Kareem Mostafa** - *Initial work* - [KareemMostafa](https://github.com/Kareemabdallah)

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

