FROM postgres

COPY createRoles.sql /docker-entrypoint-initdb.d/createRoles.sql

COPY schema.sh /docker-entrypoint-initdb.d/schema.sh

COPY schema.sql /schema.sql

