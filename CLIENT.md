# 服务端 - 管理员

## 登陆

### HTTP 请求
- POST `/v1/admin/login`
- Content-Type `application/json;charset=utf-8`

### 请求参数

```json

{
    "user": "root",
    "password": "root"
}

```

| 参数名称  | 是否必须  | 类型    | 默认值  | 描述    |
| :------- | :------- | :----- | :----- | :----- |
| user    | true     | string    |     | 只能为字母或者数字         |
| password     | true     | string    |     | 只能为字母或者数字               |


### 响应数据

```json
{
    "code": 200,
    "data": {
        "id": 1,
        "user": "root",
        "sex": 0,
        "age": 0,
        "phone": "15960611111",
        "name": "bee"
    }
}
```

| 字段名称   | 是否必须  | 类型   | 描述                              | 取值范围 |
| :---------| :------- | :----- | :-------------------------------| :------- |
| id        | true     | int    | ip黑名单ID                       |          |
| user    | true     | string    | 用户名  |          |
| sex | true     | int    | 0:男 1:女             |          |
| age        | true     | int | 年龄                          |           |
| phone     | true     | string    | 手机号                      |          |
| name     | true     | string    | 名字                      |          |



