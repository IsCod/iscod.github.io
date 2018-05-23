# HTTP协议

HTTP协议是客户机与服务器之间的请求-应答协议, 它是建立在 TCP/IP 协议上的应用层协议

## HTTP概述

HTTP是使用可靠的数据传输协议 (TCP/IP) , 从服务器上获取信息块展示给客户端的应用层传输协议

**web客户端和服务器**

web服务器是web资源的宿主

客户端常见的如web浏览器, curl也是一种客户端

### 资源、资源类型

**资源**

web资源是web内容的源头, web服务器是web资源的宿主

最简单的资源是web服务器文件系统中的静态文件, 如: 文本文件, HTML文件, JPEG图片, MP3视频文件等。资源还可以是生成动态内容的软件程序, 如PHP、Pyton等生成的内容

**资源类型**

因特网上有数千种不同的数据类型, HTTP会给每种要通过web传输的资源对象打上名为MIME类型(MIME type)的数据格式标签

web客户端从服务器获取资源后会查看MIME类型, 查看是否可以处理这类对象文件

MIME类型是一种文本标记, 表示一种主要的对像类型和一个特殊的子类型, 中间用斜杠分隔

- HTML格式的文本由 text/html 类型标记
- 普通的ASCII文本文档由 text/plain 类型标记
- JPEG格式的图片由 image/jpeg 类型标记
- GIF格式的图片由 image/gif 类型标记

常见的MIME类型有数百个, 实验性或用途有限的MIME类型则更多

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

- GET

	GET是最常用的方法, 通常用于请求服务器发送某个资源

- HEAD

	HEAD 方法和 GET 方法类似, 但服务器在响应中只返回首部, 不会返回实体的主体部分

	这就允许客户端未获取资源的情况下, 检查资源首部

	使用HEAD可以:

	```
	在不获取资源的情况下了解资源情况 (判断类型等)
	查看响应码, 看某个对象是否存在
	通过查看首部, 判断资源是否被修改
	```

- PUT 

	PUT与GET相反, PUT方法会向服务器写入文档

	PUT的语义就是让服务器用请求的的主体部分创建一个由所请求的URL命名的新文档, 因为PUT允许用户修改, 所以多少PUT需要用户用密码登录

- POST

	POST是用来向服务器输入数据的, 通常它用来支持HTML的表单	, 表单中填好的数据送到服务器, 然后服务器将其发送到它要去的地方比如一个网关程序, 然后由网关程序处理

- TRACE

	TRACE请求会在目的服务器端发起一个"环回"诊断, 行程的最后一站服务器会弹回一个TRACE响应, 并在响应主体中携带它收到的原始请求报文

- OPTIONS

	OPTIONS请求服务器告知其支持的各种功能

- DELETE

	DELETE是请求服务器删除URL所指定的资源, 但是客户端无法保证删除一定执行

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

- method (方法)

	客户端希望服务器执行的动作, 是一个单独的词如GET, POS等, 参考[方法](/TO/HTTP协议?id=方法)

- request - URL (请求URL)

	命名了所有资源，或者URL路径组件的完整URL, 参考[统一资源标识符-URI](/TO/HTTP协议?id=统一资源标识符-uri)

- version (版本)

	报文使用的HTTP版本类似

	```
	HTTP/<major>.<minor>
	```
	包含主要版本号 (major) 和次要版本号 (minor) , 切都是整数

	status-code (状态码)

	这个三位数字描述了请求过程中发生的情况, 每一个状态码的第一个数字用以描述状态的一般分类, 参考[状态码](/TO/HTTP协议?id=状态码)

- reason-phrase (原因短语)

	数字状态码的可读版本, 包含行终止序列之前的所有文本

	原因短语只对人类有意义, 因此响应行 HTTP/1.1 200 NOT OK 和 HTTP/1.1 200 OK 虽然原因短语不同，但是都应该被当做成功处理

- header (首部)

	可以有零个或多个首部组成, 每一个首部都包含一个名字, 后面跟着一个冒号 (:) 然后是一个可选的空格, 接着是一个值, 最后是一个CRLF

	首部是由一个空行的 (CRLF) 结束, 表示首部列表的结束和主体部分的开始, 参考[首部](/TO/HTTP协议?id=首部)

- entity-body (实体的主体部分)

	主体部分是一个可选的有任意数据组成的数据块

### 起始行

HTTP报文都由一个起始行开始, 请求报文的起始行说明了要做些什么, 响应报文的起始行说明发生了什么

### 首部

HTTP首部字段向请求和响应添加了一些附加信息，本质上它们是一些名/值对的列表比如Content-Type: text/plain 表示实体的主体部分是text的文本

**首部分类**

HTTP首部字段分为以下几类

- 通用首部

	即在请求和响应首部中都可以出现的首部字段

- 请求首部

	提供关于请求的信息

- 响应首部

	提供关于响应的信息

- 实体首部

	描述主体的长度和内容或者资源的首部

- 扩展首部

	规范中没有定义的新扩展字段

	每一个首部都有一个简单的语法, 名字后面跟一个冒号 (:) , 然后跟上可选的空格, 再加上字段值, 最后是一个CRLF

常见的首部实例

```
Server: nginx/1.12.2
Content-Type: text/html
Content-Length: 173
Connection: close
```

### 实体的主体部分

HTTP的第三部分是可选的实体部分, 实体的主体是HTTP报文的要传输的内容

HTTP报文可以承载多种类型的数字数据, 包含图片, 文本, 音频, HTML文档, 软件程序生成内容, 电子邮件等

### 完整的HTTP报文


请求报文实例

```
GET /index.html HTTP/1.1
Host: www.example.com

```

响应报文实例

```
HTTP/1.1 200 OK
Server: nginx/1.12.2
Date: Thu, 17 May 2018 06:38:32 GMT
Content-Type: application/json
Transfer-Encoding: chunked
Connection: keep-alive
X-Powered-By: PHP/7.0.26

{"status":200,"msg":"not found api"}

```

使用Telnet

```sh
telnet 127.0.0.1 80
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
GET / HTTP/1.1
Host: www.example.com

HTTP/1.1 200 OK
Server: nginx/1.12.2
Date: Thu, 17 May 2018 06:39:08 GMT
Content-Type: application/json
Transfer-Encoding: chunked
Connection: keep-alive
X-Powered-By: PHP/7.0.26

{"status":200,"msg":"not found api"}

```

?>注意首部和实体部分的分隔符CRLF

PHP发送请求实例:

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

PHP响应实例:

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
