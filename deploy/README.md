# 部署注意事项
总体顺序：
* 执行init.sh创建目录，创建配置文件、证书
* 创建docker-compose-env中的容器
* kafka、ES配置初始化，中间件状态检查
* 创建docker-compose中的容器

# 部署流程
将`deploy`目录上传到服务器`/tmp`下，给init.sh添加执行权限并执行
```shell
cd /tmp/deploy
chmod +x init.sh
./init.sh
```

启动中间件容器
```shell
cd /tmp/deploy/wesearch
docker-compose -f docker-compose-env.yml up -d
```
按init.sh脚本的提示及本文档后续内容完成组件初始化工作后，最后启动go服务容器
```shell
docker-compose -f docker-compose.yml up -d
```

# 组件初始化与检查
## Elasticsearch & Kibana
获取enroll token
```shell
docker exec -it elasticsearch bin/elasticsearch-create-enrollment-token --scope kibana
```
重置elasticsearch用户密码
```shell
docker exec -it elasticsearch bin/elasticsearch-reset-password -u elastic
```

使用 http://{HOST_IP}:5601 访问Kibana，使用enroll token和用户名密码登录。

参考文档：[ES官方文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)

对接完成后，可以在Kibana上创建index，获取APIKEY，然后配置到wesearch-retrieve服务的配置文件。
## 创建Index
通过Kibana创建Index。`Management -> 开发工具`
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
## 获取APIKey
`Management -> Stack Management -> API密钥`
## 停止Kibana容器
Elasticsearch初始化完成后，可以停止Kibana容器。
```shell
docker-compose -f docker-compose-env.yml stop kibana
```
## Kafka
创建Topic。进入1.kafka容器
```shell
docker exec -it kafka bash
```
执行命令：
```shell
kafka-topics.sh --create --topic create-doc --partitions 1 --replication-factor 1 --bootstrap-server 1.kafka:9092
kafka-topics.sh --create --topic parse-doc --partitions 1 --replication-factor 1 --bootstrap-server 1.kafka:9092
kafka-topics.sh --bootstrap-server 1.kafka:9092 --list
```
## Mysql
登录测试
```shell
docker exec -it mysql mysql -uroot -p123456
```
检查数据库及表初始化
```shell
use wesearch;
show tables;
select * from user;
```
## Etcd
验证返回404即可。
```shell
curl --resolve 1.etcd:2379:{容器1.etcd的IP} -i https://1.etcd:2379/v2 --noproxy "*" \
--cert /home/wesearch/certs/etcd-client.crt \
--key /home/wesearch/certs/etcd-client.key \
--cacert /home/wesearch/certs/ca.crt
```

# References
[es](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-config-bind-mount)