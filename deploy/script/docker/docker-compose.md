### yum安装pip
```shell
yum -y install epel-release

yum -y install  python-pip

//需要指定版本,直接更新最新版会出现bug
pip install --upgrade pip==20.0.1
```


### pip安装docker-compose 
```shell
pip install docker-compose
```