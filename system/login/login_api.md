## 登录相关api接口

### 1.登录接口
> 路由: [post]/api/login

> 请求参数:

|参数名|数据类型|是否必填|描述信息|
|:-:|--|--|--|
|LoginCode|string|true|账号，根据LoginType确定语义|
|Password|string|fales|密码或验证码|
|LoginType|string|fales|登录类型 psw:密码登录,qw:企业微信扫码登录,wc:微信扫码登录,min:小程序扫码登录 默认psw|
|Param|map[string]string|fales|附加参数|

> 请求示例:

```json
{
	"LoginCode": "xxx",
	"Password": "xxx"
}
```
> 返回参数:

|参数名|数据类型|描述信息|示例|
|:-:|--|--|--|
|code|int|代码|200|
|expire|string|过期时间|2022-01-08T16:06:09+08:00|
|token|string|jwt|xxxxxxxxxx|
|user|obj|用户基础信息||

> 返回示例:

```json
{
    "code": 200,
    "expire": "2022-01-08T16:06:09+08:00",
    "token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE5LCJVc2VyTmFtZSI6I",
    "user": {
        "Id": 19,
        "LoginCode": "",
        "UserCode": "1641519161",
        "UserName": "用户1641519161",
        "Phone": "15578906770",
        "Avatar": "",
        "Email": "",
        "UserType": "user",
        "Status": ""
    }
}
```

#### 1.1.账号密码
> LoginCode 支持的参数: LoginCode, UserCode, Phone, Email 任选一个
> Password 加密方式: sha256

> 请求示例:

```json
{
	"LoginCode": "xxx",
	"Password": "xxx"
}
```
#### 1.2.手机验证码
##### 1.2.1.获取手机验证码
> 路由: [get]/api/sms

> 请求参数:
|参数名|数据类型|是否必填|描述信息|
|:-:|--|--|--|
|Phone|string|true|手机号码|

> 请求示例
```
api/sms?Phone=15578906770
```

##### 1.2.2.验证码登录
> LoginCode: 手机号码
> Password:验证码

> 请求示例:
```json
{
	"LoginCode": "xxx",
	"Password": "xxx", 
    "LoginType": "sms"
}
```

#### 1.3.微信小程序
> LoginCode: code，微信授权的用户
> Param:{"phone":"手机号码"}

> 请求示例:
```json
{
	"LoginCode": "xxx",
	"Param":{"phone":"手机号码"},
	"LoginType": "mini"
}
```

#### 1.4.企业微信
> LoginCode: code，微信授权的用户

> 请求示例:
```json
{
	"LoginCode": "xxx",
	"LoginType": "work"
}
```

### 2.刷新token
> 路由: [get]/api/refresh_token
> 没有请求参数

> 返回参数:

|参数名|数据类型|描述信息|示例|
|:-:|--|--|--|
|code|int|代码|200|
|expire|string|过期时间|"2022-01-08T16:06:09+08:00"|
|token|string|jwt|xxxxxxxxxx|

### 3.注销
> 路由: [get]/api/logout
> 没有请求参数

> 返回参数:

|参数名|数据类型|描述信息|示例|
|:-:|--|--|--|
|code|int|代码|200|

### 4.获取用户登录信息
> 路由: [get]/api/login_user
> 没有请求参数

> 返回参数:

|参数名|数据类型|描述信息|示例|
|:-:|--|--|--|
|success|bool|是否成功|true|
|data|obj|用户数据||

### 5.用户注册
> 路由: [post]/api/register

> 请求参数:

|参数名|数据类型|是否必填|描述信息|
|:-:|--|--|--|
|LoginCode|string|true|账号|
|Password|string|true|密码|
|UserName|string|true|昵称|
|Phone|string|fales|手机号码|
|Email|string|fales|邮箱|

> 请求示例:
```json
{
	"LoginCode": "ceshi001",
	"Password":"12345",
	"UserName":"UserName",
	"Phone":"12345",
	"Email": "mini"
}
```

返回示例:
```json
{
    "success": true,
    "data": "注册成功",
    "total": 0,
    "current": 0,
    "pageSize": 0,
    "errorCode": "",
    "errorMessage": "",
    "showType": 0,
    "traceId": "",
    "host": "127.0.0.1"
}
```