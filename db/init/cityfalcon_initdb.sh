#!/bin/bash

# Immediately exits if any error occurs during the script
# execution. If not set, an error could occur and the
# script would continue its execution.
set -o errexit


# Creating an array that defines the environment variables
# that must be set. This can be consumed later via arrray
# variable expansion ${REQUIRED_ENV_VARS[@]}.
readonly REQUIRED_ENV_VARS=(
  "DB_USER"
  "DB_PASSWORD"
  "DB_DATABASE"
  "POSTGRES_USER")


# Main execution:
# - verifies if all environment variables are set
# - runs the SQL code to create user and database
main() {
  check_env_vars_set
  init_user_and_db
}


# Checks if all of the required environment
# variables are set. If one of them isn't,
# echoes a text explaining which one isn't
# and the name of the ones that need to be
check_env_vars_set() {
  for required_env_var in ${REQUIRED_ENV_VARS[@]}; do
    if [[ -z "${!required_env_var}" ]]; then
      echo "Error:
    Environment variable '$required_env_var' not set.
    Make sure you have the following environment variables set:
      ${REQUIRED_ENV_VARS[@]}
Aborting."
      exit 1
    fi
  done
}


# Performs the initialization in the already-started PostgreSQL
# using the preconfigured POSTGRE_USER user.
init_user_and_db() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
     CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
     CREATE DATABASE $DB_DATABASE;
     GRANT ALL PRIVILEGES ON DATABASE $DB_DATABASE TO $DB_USER;
     
EOSQL

  psql -U "$DB_USER" -d "$DB_DATABASE" <<-EOSQL
    CREATE TABLE company_info(cvr_id VARCHAR(50) PRIMARY KEY NOT NULL,created_at TIMESTAMP,address VARCHAR,postal_code_and_city VARCHAR,start_date VARCHAR,business_type VARCHAR,advertising_protection VARCHAR,status VARCHAR,telephone VARCHAR,fax VARCHAR,email VARCHAR,municipality VARCHAR,activity_code VARCHAR,secondary_names VARCHAR,financial_year VARCHAR,latest_articles_of_association VARCHAR,classes_of_shares VARCHAR,registered_capital VARCHAR,first_accounting_period VARCHAR);
EOSQL
  
}

# Executes the main routine with environment variables
main "$@"