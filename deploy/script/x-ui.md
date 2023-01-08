### 好东西
```shell
mkdir x-ui && cd x-ui
docker run -itd --network=host \
    -v $PWD/db/:/etc/x-ui/ \
    -v $PWD/cert/:/root/cert/ \
    --name x-ui --restart=unless-stopped \
    enwaiax/x-ui:latest
```

### 魔法叠加
```shell
  location /plogger {                 # plogger填你前面设置的ws的路径
          proxy_redirect off;
          proxy_pass http://127.0.0.1:3997;    # 3997换成你入站规则那边的IP
          proxy_http_version 1.1;
          proxy_set_header Upgrade $http_upgrade;
          proxy_set_header Connection "upgrade";
          proxy_set_header Host $http_host;
          proxy_read_timeout 300s;
          # Show realip in v2ray access.log
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  }

```

### 重置设置
```shell
x-ui setting reset

exit

docker restart x-ui
```