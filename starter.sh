

# 


# run postgres docker container 

docker run --env POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

# pgAdmin docker container

docker run -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4 -v pgadmin_vol:/var/lib/pgadmin -n pgadm


# Connection string for pgadmin:

# host: host.docker.internal
# port: 5432
# database: postgres
# user: postgres
# password: postgres