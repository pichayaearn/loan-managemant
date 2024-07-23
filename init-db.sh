#!/bin/bash
set -e

echo "Running init-db.sh script..."

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE loan_management;
EOSQL

echo "Database creation script completed."
