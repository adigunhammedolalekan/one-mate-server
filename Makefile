build:
	./build.sh
deploy:
	make build
	docker-compose down && docker-compose build && docker-compose up

proto:
	for d in src; do \
			for f in $$d/**/proto/*.proto; do \
				protoc --go_out=plugins=grpc:. $$f; \
				echo compiled: $$f; \
			done \
	done