postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine


createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank


migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database ""postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./... 


server :
	go run main.go


mock:
	mockgen  --package mockdb --destination db/mock/store.go  github.com/vsevicky/simplebank/db/sqlc Store

proto:
	rm -f pb/*.go 
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		proto/*.proto


.PHONY: postgres createdb dropdb migrateup migratedown migratedown1 migrateup1 sqlc test server mock proto