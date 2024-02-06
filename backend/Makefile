.PHONY: test

test:
	docker compose -f ../docker-compose.yml up -d dbtest
	cd api && go test -v -cover ./...
