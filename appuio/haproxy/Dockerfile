from haproxytech/haproxy-debian:2.7.3

RUN curl -o mysql.deb https://dev.mysql.com/get/mysql-apt-config_0.8.26-1_all.deb || exit 1 \
    dpkg -i mysql.deb || exit 1 ; \
    apt-get update || exit 1 ; \
    apt-get install -y default-mysql-client || exit 1 ; \
    rm -rf /var/lib/apt/lists/* mysql.deb ;