PASSWD := nimda
IMAGE := mysql:latest
CONTAINER := mysql


.PHONY : stop-database
stop-database:
	@docker stop $(CONTAINER)
	@docker rm $(CONTAINER)


.PHONY: start-database
start-database:
	@sudo mkdir -p /var/lib/mysql/data
	@docker run -d --name  $(CONTAINER) \
		-v /var/lib/mysql/data:/var/lib/mysql \
		-e MYSQL_ROOT_PASSWORD=$(PASSWD)\
		-p 3306:3306 $(IMAGE)


.PHONY : start-database stop-database