docker-compose-dev-up:
	docker-compose -f deployments/docker-compose.dev.yml up -d

docker-compose-dev-down:
	docker-compose -f deployments/docker-compose.dev.yml down
