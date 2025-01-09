### tar
```shell
tar -czvf test.tar.gz  data/

tar -zxvf xxxxx.tar.gz -C /specific dir 
```

### ssh
```shell
ssh -L port:127.0.0.1:port root@sship
```

### QB
```
WebUI\HostHeaderValidation=false 
WebUI\CSRFProtection=false
```

### Debain常用
```shell
sudo -i # 切换到root用户

apt update -y  # 升级packages

apt install wget curl sudo vim git -y  # 安装常用的软件

//Debian
curl -fsSL https://get.docker.com -o get-docker.sh
DRY_RUN=1 sudo sh ./get-docker.sh


sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://XXX.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker


wget -qO- yabs.sh | bash
```


### Linux查看端口占用情况
```shell
lsof -i:端口号

netstat -ntlp              //查看当前所有tcp端口
netstat -ntulp | grep 80   //查看所有80端口使用情况
```

查询CPU占用前十名
```shell
ps auxw|head -1;ps -auxf|sort -nr -k3|head -10
```

查询内存占用前十名
```shell
ps auxw|head -1;ps -auxf|sort -nr -k4|head -10
```
