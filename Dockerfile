# this dockerfile uses the ubuntu image

# bash image

FROM ubuntu

# docker user
MAINTAINER docker_user docker_user@163.com

# commands
RUN apt-get update && apt-get install -y nginx
