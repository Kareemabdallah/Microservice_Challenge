#!/bin/bash
docker tag app1:latest mostafa/app1:GoRestAPI
docker push mostafa/app1:GoRestAPI
docker tag app:latest mostafa/app:PythonRestAPI
docker push mostafa/app:PythonRestAPI