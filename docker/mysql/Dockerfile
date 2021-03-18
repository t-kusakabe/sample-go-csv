FROM mysql:5.7.20

COPY ./docker/mysql/config.d/mysql.cnf /etc/mysql/conf.d/mysql.cnf
COPY ./docker/mysql/init/init.sql /docker-entrypoint-initdb.d/init.sql
