# !/bin/sh
docker build -t rest-build -f Dockerfile1.build . 
docker run rest-build > build.tar.gz 
docker build -t rest -f Dockerfile1.dist .

docker build -t app-build -f Dockerfile2.build . 
docker run app-build > build.tar.gz 
docker build -t app -f Dockerfile2.dist .