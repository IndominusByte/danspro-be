#!/bin/bash
docker network create danspro-environment-development
echo "====== CREATE DB ======"
cd database
make dev
cd ..
echo "====== MIGRATE DB ======"
cd migration
make dev
make log-dev
cd ..
echo "====== RUNNING API ======"
cd api
make dev
make log-dev
