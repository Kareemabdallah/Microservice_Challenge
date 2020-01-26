#!/bin/sh
docker stop app1-prod
docker run --rm -p 9000:9000 --name app1-prod app1
docker stop app2-prod
docker run --rm -p 7000:7000 --name app2-prod app2