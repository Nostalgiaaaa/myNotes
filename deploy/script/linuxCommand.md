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

vim /etc/ssh/sshd_config

apt update -y  # 升级packages

apt install wget curl sudo vim git -y  # 安装常用的软件

sudo chown -R 1000:1000 ./downloads
sudo chmod -R 775 ./downloads
ls -ld ./downloads

//docker
curl -fsSL https://get.docker.com -o get-docker.sh
DRY_RUN=1 sudo sh ./get-docker.sh

//docker加速
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://XXX.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

//Storage Box
apt-get install cifs-utils
mkdir /xxx
mount.cifs -o user=你的User name,pass=你的密码 //你的CIFS地址 /xxx


//swap
swapon --show
sudo dd if=/dev/zero of=/swapfile bs=1M count=2048
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
swapon --show
echo "/swapfile none swap sw 0 0" >>/etc/fstab

//disk
fdisk -l
mkfs -t ext4 /dev/vdb
mkdir /data
mount /dev/vdb /data
df -TH
echo "/dev/vdb /data   ext4 defaults,nofail    0   0" >>/etc/fstab
mount -a

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
