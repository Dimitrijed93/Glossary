.PHONY : docker docker-build docker-run

docker: docker-db docker-build docker-run

docker-db:
	docker run --name glossaryDB -p 5432:5432 -e POSTGRES_PASSWORD=dimit -d postgres
docker-build:
	docker build -t glossary .

docker-run:
	docker run -p 9090:9090 glossary
