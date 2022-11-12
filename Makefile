up:
	docker-compose -f docker-compose.dev.yml up -d

down:
	docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans

migration:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)