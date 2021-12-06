test:
	docker-compose --file=docker-compose-test.yml down
	docker-compose --file=docker-compose-test.yml up
test-build:
	docker-compose --file=docker-compose-test.yml up --build
run:
	docker-compose up
