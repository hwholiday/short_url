# short_url
Short link 短链接服务器
### 什么是短链接   
就是把普通网址，转换成比较短的网址。比如：https://dwz.cn/XzhYJMkZ

### 原理解析
当我们在浏览器里输入 https://dwz.cn/XzhYJMkZ<br>
DNS首先解析获得 https://dwz.cn/ 的 IP 地址<br>
当 DNS 获得 IP 地址以后，会向这个地址发送 HTTP请求，查询短码 XzhYJMkZ<br>
https://dwz.cn/ 服务器会通过短码 XzhYJMkZ 获取对应的长 URL<br>
请求通过 HTTP 301 转到对应的长 URL http://www.baidu.com<br>

### 本文采用  自增序列算法 + 用户自定义短码

设置 id 自增，一个 10进制 id 对应一个 62进制的数值，1对1，也就不会出现重复的情况。<br>这个利用的就是低进制转化为高进制时，字符数会减少的特性。<br>
可使用redis Incr 实现id自增。<br>  


### 数据表设计
```base
CREATE TABLE `links` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) COLLATE utf8mb4_bin NOT NULL COMMENT '长连接',
  `keyword` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '短链接码',
  `status` tinyint(1) NOT NULL COMMENT '1系统分配 2用户自定义',
  `create_time` bigint(20) NOT NULL,
  `update_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `links_UN` (`url`),
  UNIQUE KEY `links_keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

```

### 服务使用  生成短链接
*  post   http://127.0.0.1:8099/v1/create_url <br>
*  form-data 参数   url 待转换的地址 ，status （1）自动分配  （2）用户自定义 ，keyword 自定义的短码<br>
* 请求值  url：www.google.com     status：1     keyword：
*  返回值{
    "code": 200,
    "data": "http://127.0.0.1:8099/8kBWU",
    "msg": "success"
}
