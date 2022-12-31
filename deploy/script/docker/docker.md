### 1.1.安装Docker
```shell
//官方文档
https://docs.docker.com/engine/install/centos/

yum install -y yum-utils device-mapper-persistent-data lvm2    //安装必要工具
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo  //设置yum源
yum install -y docker-ce  //下载docker
systemctl start docker   //启动docker
```

#### 1.2.安装MySQL
```shell
docker pull mysql //下载MySQL镜像
docker run --name mysql --restart=always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=密码 -d mysql //启动MySQL
```
#### 1.3.安装Redis
```shell
docker pull redis //下载Redis镜像
docker run --name redis  --restart=always -p 6379:6379 -d redis --requirepass "密码" //启动Redis
```

#### 1.4.安装nginx
```shell
docker pull nginx //下载nginx镜像

#无ssl方案
docker run --name nginx --restart=always -p 80:80 -p 443:443 -d -v /usr/local/nginx/nginx.conf:/etc/nginx/nginx.conf -v /usr/local/vue:/usr/local/vue -v /usr/local/upload:/usr/local/upload nginx  //启动nginx,映射本地配置文件

#有ssl方案
docker run --name nginx --restart=always -p 80:80 -p 443:443 -d -v /usr/local/nginx/nginx.conf:/etc/nginx/nginx.conf -v /usr/local/vue:/usr/local/vue -v /usr/local/nginx/ssl:/usr/local/nginx/ssl -v /usr/local/upload:/usr/local/upload nginx 
```


#### 1.5.安装RabbitMQ
```shell
docker pull rabbitmq:management //下载RabbitMQ镜像
docker run --name rabbit --restart=always -p 15672:15672 -p 5672:5672  -d  rabbitmq:management   //启动RabbitMQ,默认guest用户,密码也是guest。
```