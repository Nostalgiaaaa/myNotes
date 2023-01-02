### Linux查看端口占用情况
```shell
lsof -i:端口号

netstat -ntlp              //查看当前所有tcp端口
netstat -ntulp | grep 80   //查看所有80端口使用情况
```

#### tar
```shell
tar -zxvf xxxxx.tar.gz -C /specific dir 
```

### Debain常用环境
```shell
sudo -i # 切换到root用户

apt update -y  # 升级packages

apt install wget curl sudo vim git -y  # 安装常用的软件
```