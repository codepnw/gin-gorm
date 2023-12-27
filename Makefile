postgresinit:
	docker run --name gin-gorm -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 postgres:15.4

createdb:
	docker exec -it gin-gorm createdb gin-gorm