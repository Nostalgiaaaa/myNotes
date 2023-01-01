### 简易版vscode网页服务器搭建
```shell
docker run --name coder-server -p 8082:8080 -e PASSWORD=这里是你的密码 codercom/code-server:latest
```