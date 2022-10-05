#!/bin/bash

export COMPOSE_IGNORE_ORPHANS=True

# postgresql
export POSTGRESQL_IMAGE=danspro-postgresql
export POSTGRESQL_IMAGE_TAG=production
export POSTGRESQL_CONTAINER=danspro-postgresql-production
export POSTGRESQL_HOST=danspro-postgresql.service
export POSTGRESQL_USER=dansprodev
export POSTGRESQL_PASSWORD=inisecret
export POSTGRESQL_DB=danspro
export POSTGRESQL_TIME_ZONE=Asia/Kuala_Lumpur
docker build -t "$POSTGRESQL_IMAGE:$POSTGRESQL_IMAGE_TAG" -f ./manifest-docker/Dockerfile.postgresql ./manifest-docker

# redis
export REDIS_IMAGE=danspro-redis
export REDIS_IMAGE_TAG=production
export REDIS_CONTAINER=danspro-redis-production
export REDIS_HOST=danspro-redis.service
docker build -t "$REDIS_IMAGE:$REDIS_IMAGE_TAG" -f ./manifest-docker/Dockerfile.redis ./manifest-docker

# pgadmin
export PGADMIN_IMAGE=danspro-pgadmin
export PGADMIN_IMAGE_TAG=production
export PGADMIN_CONTAINER=danspro-pgadmin-production
export PGADMIN_HOST=danspro-pgadmin.service
export PGADMIN_EMAIL=admin@danspro.co.id
export PGADMIN_PASSWORD=inisecret
docker build -t "$PGADMIN_IMAGE:$PGADMIN_IMAGE_TAG" -f ./manifest-docker/Dockerfile.pgadmin ./manifest-docker

# pgbackups
export PGBACKUPS_IMAGE=danspro-pgbackups
export PGBACKUPS_IMAGE_TAG=production
export PGBACKUPS_CONTAINER=danspro-pgbackups-production
export PGBACKUPS_HOST=danspro-pgbackups.service
docker build -t "$PGBACKUPS_IMAGE:$PGBACKUPS_IMAGE_TAG" -f ./manifest-docker/Dockerfile.pgbackups ./manifest-docker

docker-compose -f ./manifest/docker-compose.production.yaml up -d --build
