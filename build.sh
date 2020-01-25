# !/bin/sh
docker build -t app1/build -f Docker .
docker run --rm -p 9000:9000 app1/build  
docker build -t app2-build -f Docker .
docker run --rm -p 5000:5000 app2/build