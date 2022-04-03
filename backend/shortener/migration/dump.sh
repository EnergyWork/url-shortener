#!/bin/sh
.\\_set_pgpass.conf.sh
D:\\Programs\\pgsql\\bin\\pg_dump -h localhost -p 5432 -U postgres -w -b -v -f url_representation.sql shortener