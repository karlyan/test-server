FROM alpine:latest

MAINTAINER Karl Yan <yanshuhuan.us@gmail.com>

RUN apk update && \
    apk add --no-cache bash openssh tzdata && \
    mkdir -p ~root/.ssh && chmod 700 ~root/.ssh/ && \
    echo -e "Port 2233\n" >> /etc/ssh/sshd_config && \
    cp /usr/share/zoneinfo/America/Chicago /etc/localtime && \
	echo "America/Chicago" > /etc/timezone && \
    cp -a /etc/ssh /etc/ssh.cache && \
    rm -rf /var/cache/apk/*

EXPOSE 2233

COPY entry.sh /entry.sh

ENTRYPOINT ["/entry.sh"]

CMD ["/usr/sbin/sshd", "-D", "-f", "/etc/ssh/sshd_config"]
