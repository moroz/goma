check-env:
	test -n "$(DATABASE_URL)" || (echo "FATAL: Environment variable DATABASE_URL is not set!"; exit 1)

db.migrate: check-env
	migrate -database "$(DATABASE_URL)" -path db/migrations up

db.rollback: check-env
	migrate -database "$(DATABASE_URL)" -path db/migrations down 1
