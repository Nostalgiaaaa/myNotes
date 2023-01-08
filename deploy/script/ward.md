#
```shell
version: '3.3'
services:
    run:
        restart: unless-stopped
        container_name: ward
        ports:
            - '4001:4000'
        environment:
            - WARD_PORT=4000
            - WARD_THEME=light
            - WARD_NAME=kiko-server
        privileged: true
        image: antonyleons/ward
```