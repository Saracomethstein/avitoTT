build:
	docker build . -t tender-service

run: build
	docker run -p 8080:8080 tender-service