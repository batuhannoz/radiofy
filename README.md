# radiofy
website where you can listen to music and talk with your friends by connecting with spotify


```bash
sudo chmod 666 /var/run/docker.sock
```

```bash
docker rm -f $(docker ps -aq) 
```

```bash
docker run --name radiofy-database \
-e MYSQL_ROOT_USER=root \
-e MYSQL_ROOT_PASSWORD=password \
-e MYSQL_DATABASE=radiofy \
-v $(pwd)/database/init.ddl:/docker-entrypoint-initdb.d/1-ddl.sql \
-p 3307:3306 \
-d mysql:8.0.23 \
--default-authentication-plugin=mysql_native_password \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_general_ci
```

```bash
APP_ENV=dev go run main.go
```

![alt text](https://github.com/batuhannoz/radiofy/blob/main/img/Club.png)
![alt text](https://github.com/batuhannoz/radiofy/blob/main/img/Radiofy.gif)
