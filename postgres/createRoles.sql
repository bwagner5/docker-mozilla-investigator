CREATE ROLE migadmin;
ALTER ROLE migadmin WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN PASSWORD 'password';

CREATE ROLE migapi;
ALTER ROLE migapi WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN PASSWORD 'password';

CREATE ROLE migscheduler;
ALTER ROLE migscheduler WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN PASSWORD 'password';

CREATE DATABASE mig OWNER migadmin;


