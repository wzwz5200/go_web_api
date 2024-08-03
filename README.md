# go_web_api
A simple golang web login API


---
title: wz
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# wz

Base URLs:

# Authentication

# 用户

## POST 添加用户VIP到期时间

POST /127.0.0.1:8080/api/add/vip

## POST 用户VIP时间查询

POST /127.0.0.1:8080/api/auth/vip

## POST 用户VIP时间查询

POST /127.0.0.1:8080/api/auth/vip


## POST 用户注册

POST /127.0.0.1:8080/api/auth/register

## POST 添加用户VIP到期时间

POST /127.0.0.1:8080/api/add/vip

>> >>>>>>>>>>>>>>>>>>>>>>>>>>>》》》》》

## POST 用户注册

POST /127.0.0.1:8080/api/auth/register


> Body 请求参数

```yaml
Names: wzwz5200
Telephone: "1111111111"
Password: 2222222222

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» Names|body|string| 是 |用户名|
|» Telephone|body|string| 是 |手机号|
|» Password|body|string| 是 |密码|

> 返回示例

> 成功

```json
{
  "code": 422,
  "data": null,
  "msg": "手机号必须为11位!"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|null|true|none||none|
|» msg|string|true|none||none|

## POST 用户登录

POST /api/auth/login

> Body 请求参数

```yaml
Names: wzwz5200
Password: xxxxxx

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|127.0.0.1/api/auth/login|query|string| 否 |none|
|body|body|object| 否 |none|
|» Names|body|string| 否 |用户名|
|» Password|body|string| 否 |用户密码|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» token|string|true|none||none|
|» msg|string|true|none||none|

## POST 用户VIP时间查询

POST /127.0.0.1:8080/api/auth/vip

> Body 请求参数

```yaml
Names: wzwz5200

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» Names|body|string| 否 |用户名|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加用户VIP到期时间

POST /127.0.0.1:8080/api/add/vip

> Body 请求参数

```yaml
Names: wzwz5200

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» Names|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型


