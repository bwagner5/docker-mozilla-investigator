#! /bin/bash 

su postgres

psql -U postgres -d mig -f schema.sql

