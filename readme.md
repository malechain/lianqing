# 接口文档

## 接口说明

所有请求都使用POST请求, Content-Type: application/json; charset=utf-8

### request
1, 接口数据都以json形式传送
```json
{
    "from": "0x0d7d912dad855a91f20d182aaaca3bc0548e1d77c93",
    "to": "0x47e9bc3504d43ef0a0132d6c6144a2fe083a3763a35",
    "value": 100
}
```
2, 数据经过RSA加密
测试私钥
```
-----BEGIN RSA Private Key-----
****
-----END RSA Private Key-----

```

测试公钥

```
-----BEGIN RSA Public Key-----
****
-----END RSA Public Key-----

```

3, RSA加密后再base64加密放入json["data"]中:

```json
{
    "data": "O9mhQVI2z5po6zZY1xinpd+UnUd9wuNihoaHhvQxsyEHpUkleiWe7aIB6ac268k5aZyNKxGJV2whoUVUG/+FZwGZ/90zorU0Pyp4cOwPU2yJ3X3+GyeWsiouw9Rq1UXV5KktcLhOJBNeyQQ3jMyPJ7rJQhdjsVTdRI9d18E3Zpi9CUO52CozA9hq5Nxb66wIRh4xK+OiUNzLi1fBcZ9aGeG8vhC5NnjmjNq6nw65GLT1aztA568Q0cJ8KMTrAc4RvNarITGoEQ55aH1fKPQluRIgtIT+XW1hmV8fZlOVsTgy8/P+4jnAtfeupEN3Ect9HA9y9nHbTqo/1IdpOrik7A=="
}
```

### response
成功实例
```json
{
    "code": 200,
    "message": "提交成功!",
    "data": "enmy05qGb1f0pf/gO9eHE0nxLWsrXmOIpJ7cHewPEJ8UFYdWm3cvE6I4xFNnvvqsfmtG9TjYCETRhgq/IpqXA8yvtLfBCoxdmhUt8gRI+T81w8t/4UYR7yzSojiTHmR3Vkfu4TBwgH8F6IbWDU0m415rL26iQXSIJNXMTJBe4Pp/pgFF+AHydQVeysZlHiabvTE7UrLq1EZV2pNPk0wPFXWyaGBCOTNKIN8Zba/V+bs+CW6sRhZVqdAvZZxmTnYgmK50c7yPlrXRB9fNkYUpUpGrhWhJy3HYhTWz1VqD9jv88AWi2iB9+staVGojIEg8/+W3ChHnafQN7hA16cJgPw=="
}
```

错误实例
```json
{
    "code": 400,
    "message": "base64解密出错",
    "data": null
}
```

data 先base64解密后RSA解密得到json格式

## 接口
1, 获取一个新的地址

接口地址:
/v1/createAccount

request
```
无
```

response
```json
{
    "code": 200,
    "message": "OK",
    "data": "XnJksC9RiEReEQmgA1i07LHPqpPSBATJa0bxUZA96kKwdOcNQLW5FxM9mdA8nKKHIJQmhOjaG89Z3Z7sAzi4LWJ1DM0jP08T9ECqadsHGUZGMI7uSn4aiJHFWxKez/PT3X92NiMjV3wwNBcdMgwipgvRCpWzXbdP+SxFpU1/vcqr5dEk4MZrO64OZrD/BEL/kkbwNEWXCqNE1v7GiWVHzj0YE1kmgoqLrsyDTZ6xfkWGvCpo52GdP0WButoXM0Uut4l5V9K4vnVMnxYHu4siUKIiSVQ0z0JbEuEUHhiDdAhs2gYWu/r5sv84pQE8uV13wVjwHZUU/NbV2bKDs8nMtA=="
}
```
data解密后
```json
{
    "addr": "0xd769CAB1108cDDaF35A4480a5f8258754E948419"
}
```

2, 提币接口

接口地址:
/v1/extract

request
```json
{
    "from": "0x0d7d912dad855a91f0d182aaaca3bc0548e1d77c",
    "to": "0x47e9bc3504d43ef0a032d6c6144a2fe083a3763a",
    "value": 100
}
```
response
```json
{
    "code": 200,
    "message": "转账成功!",
    "data": "R4eW5ANc9yjTSQzcnxT9WRneQMAVnpR1jWGy1JsNfKb+BSclSYQKS2Xl1zAewVqsFAPdYnUrE1Na8IqmFgc5nnLnG/t/wTx4ng3EC4555M3rRdA/Mt7uZ145gaor8guMQZaQhd+8370KjoGvkkDkWavfDl5kWmk7dfjmbizjgZ8wDIcfvhbe1pATgki9u/7AWpsBpMi+ykMwATIe/hbvsq6aSy460x1EEgWnF2ZqvfgP7Q06aawyuF8YUm+LKZy2GxAd/jNUKX94wAfWlDIOWc7IgU95B7r80D6czPbmmnEU1WuAHIMF6as5p3+f+0XDZ+Ay0m7zo7dTKHT2ppQt8A=="
}
```
data解密后 得到交易hash
```json
{
    "txid":"0xb09658259b74302a9d64d237b6005541afc96db60a80d61f02ab15994ca81188"
}
```
