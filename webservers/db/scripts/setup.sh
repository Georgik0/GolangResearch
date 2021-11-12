#psql -U db_user -d app_db <<-EOSQL
#  CREATE TABLE test_table (
#      first_column text
#  );
#EOSQL

psql -U db_user -d app_db < docker-entrypoint-initdb.d/init.sql