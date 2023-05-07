# 部署注意事项
总体顺序：
* 创建目录，创建配置文件、证书
* 创建docker-compose-env中的容器
* kafka、ES配置初始化
* 创建docker-compose中的容器

# 宿主机公共目录创建
创建目录并上传配置文件，证书文件。
```shell
mkdir -p /home/wesearch/{certs,etc,yml}
```
# Mysql
目录创建
```shell
mkdir -p /home/wesearch/mysql/{conf,data,log,init}
```

拷贝初始化sql
```shell
cp *.sql /home/wesearch/mysql/init
```

写入配置文件
```shell
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
```
登录测试
```shell
docker exec -it mysql mysql -uroot -p123456
```
# Elasticsearch & Kibana
## 部署
ES和Kibana各自都有自己的config,data,logs,plugins目录，需要将这些目录挂载在宿主机。

由于启动时，这两个组件都需要读取自己的配置文件，可以在启动容器前在宿主机目录上创建好配置文件，填写必要的配置。
```shell
mkdir -p /home/wesearch/elasticsearch/{config,data,logs,plugins}
mkdir -p /home/wesearch/kibana/{config,data,logs,plugins}

mv es/* /home/wesearch/elasticsearch/config

chown -R 1000:0 /home/wesearch/elasticsearch
chown -R 1000:0 /home/wesearch/kibana

cat > /home/wesearch/elasticsearch/config/elasticsearch.yml <<EOF
network.host: 0.0.0.0
cluster.name: "docker-cluster"
EOF

cat > /home/wesearch/kibana/config/kibana.yml <<EOF
server.host: 0.0.0.0
i18n.locale: "zh-CN"
EOF

```

ES初次启动后，会在控制台打印enroll token，用于在Kibana初次登录页面配置对接ES。 

参考文档：


对接完成后，可以在Kibana上获取APIKEY，然后配置到wesearch-retrieve服务的配置文件。
## 创建Index
通过接口或Kibana创建Index
```shell
PUT wesearch
PUT wesearch/_mapping
{
    "_source": {
      "enabled": false
    },
    "properties": {
      "content": {
        "type": "text",
        "store":false
      }
    }
}
GET wesearch
```
# Kafka
## 宿主机目录创建
```shell
mkdir -p /home/wesearch/kafka/data
```
## 创建Topic
进入1.kafka容器，执行命令：
```shell
kafka-topics.sh --create --topic create-doc --partitions 1 --replication-factor 1 --bootstrap-server 1.kafka:9092
kafka-topics.sh --create --topic parse-doc --partitions 1 --replication-factor 1 --bootstrap-server 1.kafka:9092
kafka-topics.sh --bootstrap-server 1.kafka:9092 --list
```

# Etcd
## 宿主机目录创建
```shell
mkdir -p /home/wesearch/etcd/data
```

# References
[es](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-config-bind-mount)