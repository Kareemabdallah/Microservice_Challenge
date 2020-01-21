#!/bin/sh
docker stop app1
docker run --rm -p 9000:9000 --name app1 rest
docker stop app2
docker run --rm -p 5000:5000 --name app2 app2