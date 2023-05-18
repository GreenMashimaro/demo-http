# Go Http Request Demo

## chapter3_2

```
curl https://www.baidu.com
```

## chapter3_4

```
curl --get --data-urlencode "query=hello world" http://www.baidu.com
```

## chapter3_6

```
curl --head http://www.baidu.com
```

## chapter3_7

x-www-form-urlencoded


```
curl -d test=value http://www.baidu.com
```

## chapter3_8

multipart 表单


```
curl -T data.txt -H "Content-Type: text/plain" http://www.baidu.com
```

## chapter3_8_2

like chapter3_8, the difference is using `strings.NewReader`

## chapter3_9

multipart/form-data

default Content-Type = `application/octet-stream`

```
curl -F "name=Michael Jackson" -F "thumbnail=@photo.png" http://www.baidu.com
```

## chapter3_11

like chapter3_9, the difference is set `Content-Type`

## chapter3_15

```
curl -x http://www.baidu.com http://github.com
```

## chapter3_19

```
curl -X DELETE http://www.baidu.com
```

## chapter3_20

```
curl -m 2 http://www.baidu.com/slow_page
```
