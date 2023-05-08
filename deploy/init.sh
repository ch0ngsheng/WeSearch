#!/bin/bash

# run as root as /tmp/deploy dir.

echo "common"
mkdir -p /home/wesearch
cp -r /tmp/deploy/certs /home/wesearch
cp -r /tmp/deploy/etc /home/wesearch

echo "mysql"
mkdir -p /home/wesearch/mysql/{conf,data,log,init}
cp sql/*.sql /home/wesearch/mysql/init/

cat > /home/wesearch/mysql/conf/my.cnf <<EOF
[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
lower_case_table_names=1 # 不区分大小写
# By default we only accept connections from localhost
#bind-address   = 127.0.0.1
default-time_zone = '+8:00'

symbolic-links=0
character-set-server=utf8mb4
[client]
default-character-set=utf8mb4
[mysql]
default-character-set=utf8mb4

EOF

echo "kafka"
mkdir -p /home/wesearch/kafka/data

echo "etcd"
mkdir -p /home/wesearch/etcd/data
chown -R 1001:1001 /home/wesearch/etcd

echo "elasticsearch & kibana"
mkdir -p /home/wesearch/elasticsearch/{config,data,logs,plugins}
mkdir -p /home/wesearch/kibana/{config,data,logs,plugins}

cp /tmp/deploy/es/* /home/wesearch/elasticsearch/config
touch /home/wesearch/elasticsearch/config/users
touch /home/wesearch/elasticsearch/config/users_roles

cat > /home/wesearch/elasticsearch/config/elasticsearch.yml <<EOF
network.host: 0.0.0.0
cluster.name: "docker-cluster"
EOF

cat > /home/wesearch/kibana/config/kibana.yml <<EOF
server.host: 0.0.0.0
i18n.locale: "zh-CN"
EOF

chown -R 1000:0 /home/wesearch/elasticsearch
chown -R 1000:0 /home/wesearch/kibana

echo -e "\n\nTIPS: after deploy containers, should:"
echo -e "1. connect es using kibana, generate APIKey and write to etc/retrieve.yaml
2. create es index
3. create kafka topic
4. configure etc/*.yaml
5. check mysql init result"