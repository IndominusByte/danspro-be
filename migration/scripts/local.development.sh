#!/bin/bash

export COMPOSE_IGNORE_ORPHANS=True

# db migration
export DB_MIGRATION_IMAGE=danspro-db-migration
export DB_MIGRATION_IMAGE_TAG=development
export DB_MIGRATION_CONTAINER=danspro-db-migration-development
export DB_MIGRATION_HOST=danspro-db-migration.service
export BACKEND_STAGE=development

docker build -t "$DB_MIGRATION_IMAGE:$DB_MIGRATION_IMAGE_TAG" .
docker-compose -f ./manifest/docker-compose.development.yaml up -d --build

