# Makefil
# Dockerfile不支持父目录拷贝，因此直接在父目录运行docker build

# 设置变量
CLIENT_IMAGE_NAME = greeter_client_image
SERVER_IMAGE_NAME = greeter_server_image
VERSION = 1.0
CLIENT_CONTAINER_NAME = greeter_client_container
SERVER_CONTAINER_NAME = greeter_server_container

# 构建 Docker 镜像
build:
	docker build -t $(CLIENT_IMAGE_NAME):$(VERSION) -f greeter_client/Dockerfile .
	docker build -t $(SERVER_IMAGE_NAME):$(VERSION) -f greeter_server/Dockerfile .

buildc:
	docker build -t $(CLIENT_IMAGE_NAME):$(VERSION) -f greeter_client/Dockerfile .
	
builds:
	docker build -t $(SERVER_IMAGE_NAME):$(VERSION) -f greeter_server/Dockerfile .

# 运行容器
# -p host_ip:container_ip
runs:
	docker run -d -p 50051:50051 --name $(SERVER_CONTAINER_NAME) $(SERVER_IMAGE_NAME):$(VERSION)
	
runc:
	docker run -d --name $(CLIENT_CONTAINER_NAME) $(CLIENT_IMAGE_NAME):$(VERSION)

logc:
	docker logs $(CLIENT_CONTAINER_NAME)

logs:
	docker logs $(SERVER_CONTAINER_NAME)
	
# 停止容器
stop:
	docker stop $(SERVER_CONTAINER_NAME) $(CLIENT_CONTAINER_NAME)

stops:
	docker stop $(SERVER_CONTAINER_NAME)
	
stopc:
	docker stop $(CLIENT_CONTAINER_NAME)
	
# 移除容器
remove:
	docker rm $(SERVER_CONTAINER_NAME) $(CLIENT_CONTAINER_NAME)
	
removes:
	docker remove $(SERVER_CONTAINER_NAME)
	
removec:
	docker remove $(CLIENT_CONTAINER_NAME)

# 清理：停止并移除容器，然后删除镜像
clean: stop remove

cleani:
	docker rmi $(CLIENT_IMAGE_NAME):$(VERSION) $(SERVER_IMAGE_NAME):$(VERSION)

