build:
	go build -v -o bin/migration *.go

run: build
	bin/migration

dev:
	bash ./scripts/local.development.sh

prod:
	bash ./scripts/local.production.sh

log-dev:
	docker logs -f --tail 10 danspro-db-migration-development

log-prod:
	docker logs -f --tail 10 danspro-db-migration-production
