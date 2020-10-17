# redis

## docker

```bash
docker run -d --name redis --net redis -p 6379:6379 -d redis:alpine
```

## 驱动

```bash
go get -u github.com/go-redis/redis/v8
```
