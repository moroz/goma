guard-%:
	@ test -n "${$*}" || (echo "FATAL: Environment variable $* is not set!"; exit 1)

db.test.prepare: guard-TEST_DATABASE_NAME guard-TEST_DATABASE_URL
	createdb ${TEST_DATABASE_NAME} || true
	@ env GOOSE_DBSTRING="${TEST_DATABASE_URL}" goose up
