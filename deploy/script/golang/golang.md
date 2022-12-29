```shell

cd /usr/local

#下载安装包
wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz

#解压安装包
tar -zxvf go1.19.1.linux-amd64.tar.gz


vim /etc/profile

#添加配置
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin


#保存退出
:wq
 
#配置生效
source /etc/profile

```

执行 go version 命令查看Go版本如下，则已安装成功。

![image.png](https://kikoblog.oss-cn-hongkong.aliyuncs.com/articles/79d4527eb0aab12617e736f2ba0e84d5.png)