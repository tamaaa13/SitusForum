export $(shell xargs < .env)

export MYSQL_URL='mysql://root:$(MYSQL_ROOT_PASSWORD)@tcp(localhost:$(DB_PORT))/$(MYSQL_DATABASE)'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations	up

migrate-down:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations down
