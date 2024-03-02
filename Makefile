guard-%:
	@ test -n "${${*}}" || (echo "Environment variable ${*} is not set!"; exit 1)

db.migrate: guard-DATABASE_URL
	migrate -database "$(DATABASE_URL)" -path db/migrations up

db.rollback: guard-DATABASE_URL
	migrate -database "$(DATABASE_URL)" -path db/migrations down 1

