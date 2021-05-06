postgres:
		docker run --name qf_postgres -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdbtest:
		docker exec -it qf_postgres createdb --username=root --owner=root qf_test

createdbdev:
		docker exec -it qf_postgres createdb --username=root --owner=root qf_dev

dropdbtest:
		docker exec -it qf_postgres dropdb qf_test

dropdbdev:
		docker exec -it qf_postgres dropdb qf_dev

migrateuptest:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_test?sslmode=disable" -verbose up 

migratedowntest:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_test?sslmode=disable" -verbose down

migrateup1test:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_test?sslmode=disable" -verbose up 1

migratedown1test:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_test?sslmode=disable" -verbose down 1


migrateupdev:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_dev?sslmode=disable" -verbose up 

migratedowndev:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_dev?sslmode=disable" -verbose down

migrateup1dev:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_dev?sslmode=disable" -verbose up 1

migratedown1dev:
		migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5430/qf_dev?sslmode=disable" -verbose down 1

test:
		go test -v -cover ./...

sqlc: 
	sqlc generate

.PHONY: postgres createdbtest dropdbtest createdbdev dropdbdev migrateuptest migratedowntest migrateup1test migratedown1test migrateupdev migratedowndev migrateup1dev migratedown1dev sqlc
