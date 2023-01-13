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

### Debain常用
```shell
sudo -i # 切换到root用户

apt update -y  # 升级packages

apt install wget curl sudo vim git -y  # 安装常用的软件

#体检
wget -N --no-check-certificate https://raw.githubusercontent.com/V2RaySSR/vps/master/vpstest.sh && bash vpstest.sh

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
