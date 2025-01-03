### Linux查看端口占用情况
```shell
lsof -i:端口号

netstat -ntlp              //查看当前所有tcp端口
netstat -ntulp | grep 80   //查看所有80端口使用情况
```

### tar
```shell
tar -czvf test.zip.tar  test/

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

#体检
curl -LsO bench.monster/speedtest.sh; bash speedtest.sh -Asia
wget -q https://github.com/Aniverse/A/raw/i/a && bash a
wget -qO- yabs.sh | bash

#综合工具箱
wget -O box.sh https://raw.githubusercontent.com/BlueSkyXN/SKY-BOX/main/box.sh && chmod +x box.sh && clear && ./box.sh

#回程
curl https://raw.githubusercontent.com/zhucaidan/mtr_trace/main/mtr_trace.sh|bash  # 检测回程脚本

#奈飞脚本
bash <(curl -L -s https://raw.githubusercontent.com/lmc999/RegionRestrictionCheck/main/check.sh)
```


查询CPU占用前十名
```shell
ps auxw|head -1;ps -auxf|sort -nr -k3|head -10
```

查询内存占用前十名
```shell
ps auxw|head -1;ps -auxf|sort -nr -k4|head -10
```
