up ::
	docker-compose -f docker-compose.dev.yml up -d

down ::
	docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans

create_db ::
	docker exec -it arc_bank_system_postgres_1 createdb --username=root --owner=root arc_bank

drop_db ::
	docker exec -it arc_bank_system_postgres_1 dropdb arc_bank

migration ::
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate_up ::
	migrate -path internal/db/migrations --database "postgresql://root:root@localhost:5434/arc_bank?sslmode=disable" -verbose up

migrate_down ::
	migrate -path internal/db/migrations --database "postgresql://root:root@localhost:5434/arc_bank?sslmode=disable" -verbose down

generate_orm ::
	sqlc generate