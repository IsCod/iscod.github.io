# HTTP协议

HTTP协议是客户机与服务器之间的请求-应答协议, 它是建立在 TCP/IP 协议上的应用层协议

## HTTP概述

HTTP是使用可靠的数据传输协议 (TCP/IP) , 从服务器上获取信息块展示给客户端的应用层传输协议

**web客户端和服务器**

web服务器是web资源的宿主

客户端常见的如web浏览器, curl也是一种客户端

### 资源、资源类型

**资源**

web服务器是web资源的宿主, web资源是web内容的源头

最简单的资源是web服务器文件系统中的静态文件如:文本文件, HTML文件, JPEG图片, MP3视频文件等。资源还包含生成动态内容的软件程序如PHP、Pyton等生成的内容

**资源类型**

因特网上有数千种不同的数据类型, HTTP会给每种要通过web传输的资源对象打上名为MIME类型(MIME type)的数据格式标签

web客户端从服务器获取资源后会查看MIME类型, 查看是否可以处理这类对象文件

MIME类型是一种文本标记, 表示一种主要的对像类型和一个特殊的子类型, 中间用斜杠分隔

- HTML格式的文本由 text/html 类型标记
- 普通的ASCII文本文档由 text/plain 类型标记
- JPEG格式的图片由 image/jpeg 类型标记
- GIF格式的图片由 image/gif 类型标记

### 统一资源标识符 (URI) 

每一个服务器资源都有一个名字, 这样客户端就可以根据名字访问该资源了

统一资源标识符 (URI) 就是服务器资源名简称, 它用于在世界范围内唯一标示并定位信息资源

URI有两种形式分别为:统一资源定位符 (URL) 和统一资源名 (URN)

**统一资源定位符(URL)**

URL 是资源标示最常见的形式, 大部分的URL是一种标准格式, 这种格式包含三部分

- 第一部分称为方案, 说明协议的类型, 最常见的就是HTTP协议 (http://)
- 第二部分给出服务的因特网地址 (比如: www.example.com)
- 其余部分指定了web服务器上的某个资源 (比如: /image/my-header.jpg)

**例如:**

```
http://www.example.com/image/my-header.jpg
```

**统一资源名(URN)**

URN 是作为特定资源的唯一名称使用, 与当前资源的位置无关

使用与位置无关的URN, 就可以将资源位置随意改动. 但是URN仍然处于试验阶段, 尚未大范围使用

**例如:**

```
urn:iscod:rfa:1234
```

?> 目前几乎所有的URI都是URL

### 事务

一个完整的HTTP事务由一条从客户端发起的请求命令和一个服务器发回客户端的相应结果组成

这种通信通过名为HTTP报文的格式化数据块进行交流

### 方法

HTTP支持不同的请求命令, 这些命令称为HTTP方法 (method)

每一条HTTP请求报文都包含一个方法, 这个方法告诉服务器需要执行什么样的动作 (获取, 删除, 还是更新等)

**常见方法**

HTTP Method | 描述
----------- | -----------------------------
GET         | 从服务器向客户端发送资源
POST        | 将客户端数据发送到服务器应用程序
PUT         | 将客户端资源存储到服务器资源中
DELETE      | 从服务器中删除资源

### 状态码

每条HTTP响应报文都会携带一个状态码, 告知客户端请求是否成功

响应吗是一个三位数字的代码, 伴随着响应码, HTTP还会发送一条解释行的描述

HTTP状态码分为5大类, 虽然没有实际的规范对解析描述度短语进行规范, 但是HTTP/1.1还是推荐规范使用原因短语

**状态码分类**

```
100 ~ 199 信息性状态码
200 ~ 299 成功状态码
300 ~ 399 重定向状态码
400 ~ 499 客户端错误状态码
500 ~ 599 服务器错误状态码
```

**常见响应码和描述**

```
200 OK
404 Not Found
502 Bad Gateway
```

### 连接

HTTP是一个应用层协议, 这样HTTP就不用关心网络通信的具体细节, 它将联网通信的细节都交给通用可靠的TCP/IP协议进行管理

TCP提供以下特性

- 无差错的数据传输
- 按序传输 (数据总是会按照发送的顺序到达)
- 未分段的数据流 (可以在任意时刻以任意尺寸发送数据)

只要建立了TCP连接, 客户端与服务端之间的报文交换就不会丢失, 不会被破坏, 也不会出现错序

**连接、IP地址和端口**

HTTP客户端向服务器发送报文之前, 需要在网际协议间通过IP地址和端口建立客户端与服务端的TCP/IP连接

客户端如何通过URL得知服务器的IP和端口?

我们看以下几种URL格式

- http://127.0.0.1:80
- http://www.example.com:80/index.html
- http://www.example.com/index.html

第一种URL使用了服务器IP和端口即:127.0.0.1和80

第二种URL没有数字形式的IP地址, 它使用了一种文本形式的域名或者称为主机名 (www.example.com)
主机名是人性的化的IP地址别称, 可以通过域名解析服务 (DNS) 的机制查询出主机名对应的IP地址

第三种和第二种类似但是没有端口号, 因为URL中默认端口号是80, 这样客户端就可以方便的建立到服务器的TCP连接了

**浏览器客户端查询IP和端口获取资源过程**

- 浏览器从 URL 中解析出服务器的主机名
- 浏览器将服务器的主机名转换成服务器的 IP 地址
- 浏览器将端口号(如果有的话)从 URL 中解析出来
- 浏览器建立一条与 Web 服务器的 TCP 连接
- 浏览器向服务器发送一条 HTTP 请求报文
- 服务器向浏览器回送一条 HTTP 响应报文
- 关闭连接,浏览器显示文档

**使用telnet**

Telnet程序可以连接到某个目标的TCP端口, 并将此TCP端口的回送显示到用户屏幕上, 它可以连接技术所有的TCP服务器。

可以通过Telnet与web服务器进行通话。

```sh
$ telnet 127.0.0.1 80
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

GET /index.html HTTP/1.1
Host: www.example.com

HTTP/1.1 200 OK
Server: nginx/1.12.2
Date: Wed, 16 May 2018 11:44:29 GMT
Content-Type: text/html
Content-Length: 173
Connection: close

<html>
<\html>
```

Telnet可以很好的模拟HTTP客户端, 但是不能作为服务器使用, 但是对Telnet做脚本自动化是非常繁杂的

如果需要更灵活的工具可以尝试nc(netcat), nc可以方便的操作基于UDP和TCP的流量

## HTTP报文

HTTP报文是在HTTP应用程序之间发送的数据块, 这些数据以一些文本形式的元信息开头, 后面跟上可选的数据部分组成

这些报文在客户端、服务器, 代理之间流动。"流入", "流出", "上游", "下游"是用来描述报文方向的术语。

### 报文流

### 报文的组成部分

HTTP报文是简单的格式化数据块, 每条报文都包含一条来自客户端的请求, 和一条服务端的响应

它们由三部分组成

- 对报文进行描述的起始行
- 包含属性的首部块
- 可选的, 包含数据主体的body部分

**起始行**和**首部**都是由行分隔的ASCII文本

每行都以一个由两个字符组成的行终止序列作为结束, 其中包括一个回车符号 (ASCII吗13) 和一个换行符号 (ASCII吗10), 这个终止序列称为**CRLF**

主体部分是一个可选的数据块, 与起始行和首部不同的是可以包含文本或者二进制数据

### 报文的语法

所有的HTTP报文都可以分为**请求报文**和**响应报文**两类

请求报文会向服务器请求一个动作, 响应报文会将请求结果返回给客户端, 请求报文和响应报文的基本报文结构相同

**请求报文的格式**

```
<method> <request - URL> <version>
<headers>
<entity-body>
```

**响应报文的格式**

```
<version> <status-code> <reason-phrase>
<headers>
<entity-body>
```

?> 响应报文与请求报文只有起始行不同

- method ([方法](/HTTP协议?id=方法)

客户端希望服务器执行的动作, 是一个单独的词如GET, POST

- request - URL (请求URL)

命名了所有资源，或者URL路径组件的完整URL

- version (版本)

报文使用的HTTP版本类似

```
HTTP/<major>.<minor>
```
包含主要版本号 (major) 和次要版本号 (minor) , 切都是整数

status-code (状态码)

这个三位数字描述了请求过程中发生的情况, 每一个状态码的第一个数字用以描述状态的一般分类参考[状态码](/HTTP协议?id=状态码)

- reason-phrase (原因短语)

数字状态码的可读版本, 包含行终止序列之前的所有文本

原因短语只对人类有意义, 因此响应行 HTTP/1.1 200 NOT OK 和 HTTP/1.1 200 OK 虽然原因短语不同，但是都应该被当做成功处理

- header (首部)

可以有零个或多个首部组成, 每一个首部都包含一个名字, 后面跟着一个冒号 (:) 然后是一个可选的空格, 接着是一个值, 最后是一个CRLF

首部是由一个空行的 (CRLF) 结束, 表示首部列表的结束和主体部分的开始

- entity-body (实体的主体部分)

主体部分是一个可选的有任意数据组成的数据块

### 起始行

HTTP响应起始行有三部分组成, 协议HTTP版本、状态码、描述发送状态行

### 首部

HTTP报文包含请求和响应报文两部分, 请求和响应都由: 起始行、首部字段、主体三部分组成

**起始行**

报文的第一行就是起始行, 在请求报文中主要用来说明做些什么, 在响应报文中说明出现了什么情况

**首部字段**

起始行后面有零个或者多个首部字段。每个首部字段都包含一个名字和一个值, 为了便与解析, 两者之间用冒号 (:) 分割。首部字段以空行结束。

**主体**

空行之后就是可选的主体部分, 其中包含了所有类型的数据。请求报文主体中包含要发送给Web服务器的数据, 响应主体中装载了要返回给客户端的数据, 主体可以包含任意的二进制数据 (比如图片, 视频, 音频等) , 当然文本是主体中最常见的。

### 请求报文

HTTP请求的起始行有方法, 请求地址, 和协议版本字段组成

**方法**
HTTP协议的方法包含POST,GET,PUT,DELETE等方法

**例如:**

```
POST /index.php HTTP/1.0
```

HTTP的请求首部字段包含一个名字和一个值, 中间又冒号(:)分割

常见的请求首部字段包含User-Agent、HOST、Accept、Accept-language、Cache-Control、Connection、Date、Pragma、Transfer-Encoding、Upgrade、Via等

**例如:**

```
Host: example.com
User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
```

**PHP实例:**

```php
<?php
	$headers = array(
		'POST /index.php HTTP/1.0',
		'User-Agent:CMD (Linux:Intel Linux)',
		'HOST: example.com',
		'Agent: text/*',
		'Agent-language: en'
		'Content-type: multipart/form-data;charset="utf-8"',
		'Cache-Control: no-cache',
		'Content-length:100',
	);

	$ch = curl_init();

	curl_setopt($ch, CURLOPT_URL, 'example.com');

	curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);//文本返回结果

	curl_setopt($ch, CURLOPT_TIMEOUT, 10);//执行最长秒数

	curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

	$data = curl_exec($ch);

	if (curl_errno($ch)) {
		echo "Error: " . curl_error($ch);
	} else {
		var_dump($data);
		curl_close($ch);
	}
?>
```

**HTTP响应报文**

HTTP响应起始行有三部分组成, 协议HTTP版本、状态码、描述发送状态行

**例如:**

```
HTTP/1.O 404 Not Found
```

HTTP响应首部字段主要是不能放到状态行的附加信息

**例如:**

```
Server: Nginx/1.6.3
Content-length: 403
Content-type: text/html
Date:Mon, 05 Sep 2016 07:05:22 GMT
```

**PHP实例:**

```php
<?php
    header('HTTP/1.0 200 Ok');
    header('Cache-Control: no-cache, must-revalidate');//无缓存
    header('Content-Type: image/png');
    header('Expires: ' . gmdate(DATE_RFC822, time()-3600*24));//过期时间
?>
```

**如何自定义头域和POST数据并获取？**

头域的信息验证在很多的API调用中用到, 如下头域中添加了user-appKey和user-appSecret, 并POST一组数据

**PHP实例:**

```php
<?php
	$headers = array(
	    'POST /index.php HTTP/1.0',//该行不影响请求, 可删除
	    'Content-Type: multipart/form-data;charset="utf-8"',//注意这里的参数设置
	    'User-AppKey: youAppKey',//自定义的头域
	    'User-AppSecret: youAppSercret'
	);

	$post_data = array(
	    'uid' => '10000',
	    'nickName' => '李华',
	);

	/*
	*  如果Content-Type设置为application/x-www-form-urlencoded, 需$post_data以urlencode形式
	*  $o = '';
	*  foreach ($post_data as $key => $value){
	*    $o .= $key . '=' . urlencode($value) . '&';
	*  }
	*
	*  $post_data = substr($o, 0, -1);
	*/

	$ch = $curl_init();

	curl_setopt($ch, CURLOPT_URL, 'example.com/index.php');

	curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

	curl_setopt($ch, CURLOPT_POST, TRUE);

	curl_setopt($ch, CURLOPT_POSTFIELDS, $post_data);

	$data = curl_exec($ch);

	if ($curl_errno($ch)) {
	    echo 'Error: ' . curl_error($ch);
	} else {
	    var_dump($data);
	    curl_close($ch);
	}
?>
```

**PHP获取HTTP请求信息**

使用EGPCS标识获取即:
```
$_COOKIE、$_GET、$_POST、$_FILES、$_SERVER、$_ENV
```

**PHP实例:**

```php
<?php
	$appKey = $_SERVER['HTTP_USER_APPKEY'];
	$appSecret = $_SERVER['HTTP_USER_APPSECRET'];
	$nickName = $_POST['nickName'];
	echo 'AppKey: ' . $appKey . "\n";
	echo 'AppSecret: ' . $appSecret . '\n';
	echo 'nickName: ' . $nickName;
?>
```

**输出结果:**

```
AppKey: youAppKey
AppSecret: youAppSercret
nickName: 李华
```


## 连接管理