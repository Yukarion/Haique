test:
	rm -rf build/redis/data/test/*
	docker-compose --file=docker-compose-test.yml up
test-build:
	docker-compose --file=docker-compose-test.yml up --build
run:
	docker-compose up
run-build:
	docker-compose up --build
