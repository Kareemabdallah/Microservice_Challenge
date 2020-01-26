# !/bin/sh

# Timestamp Function
timestamp() {
	date +"%T"
}

# Go build
echo "🐋	$(timestamp): building image app1:test"
    docker build -t app1:test -f Dockerfile .
    docker run --rm -p 9000:9000 app1:test

    docker build -t app2:test -f Dockerfile2 .
    docker run --rm -p 7000:7000 app2:test
    

echo "🌧️	 $(timestamp): deploying to Minikube"
	kubectl delete deployment app1
	kubectl delete service app1
	kubectl apply -f deploy.yml