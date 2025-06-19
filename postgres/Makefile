DC := docker-compose

up: 
	$(DC) up -d 

down: 
	$(DC) down 

logs:
	$(DC) logs -f

psql:
	PGPASSWORD=mypassword psql -h localhost -p 5432 -U myuser -d mydb
