APP_SRC_GOS = ${shell find ./app -name '*.proto'}
APP = ./build/app

.PHONY: app clean

app: $(APP)

$(APP): $(APP_SRC_GOS)
	docker-compose run --rm build-container bash -c "cd /srv/app && GOOS=linux GOARCH=amd64 go build -o ./../${APP}"

image:
	cd build && docker build -f ../docker/Dockerfile -t g-actions-test .

clean:
	rm build/*