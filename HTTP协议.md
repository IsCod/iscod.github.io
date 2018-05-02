# HTTP协议

HTTP协议是客户机与服务器之间的请求-应答协议

## 资源、资源类型

web资源是web内容的源头，最简单的资源是web服务器文件系统中的静态文件如：文本文件，HTML文件，JPEG图片，MP3视频文件等。资源还包含生成动态内容的软件程序如PHP、Pyton等生成的内容。

HTTP对通过Web传输的对象打上名为：MIME类型的数据标签。MIME类型是一种文本标记，表示一种主要的对像类型和一个特殊的子类型，中间用斜杠分隔。

- HTML格式的文本由 text/html 类型标记
- 普通的ASCII文本文档由 text/plain 类型标记
- JPEG格式的图片由 image/jpeg 类型标记
- GIF格式的图片由 image/gif 类型标记

## 统一资源标识符（URI）

服务器资源名被称为统一资源标识符（URI）,用于在世界范围内唯一标示并定位信息资源，URI有两种形式分别为：URL和URN。

**统一资源定位符（URL）**是资源标示最常见的形式，大部分的URL是一种标准格式，这种格式包含三部分

- 第一部分称为方案，说明协议的类型，最常见的就是HTTP协议(http://)
- 第二部分给出服务的因特网地址（比如：www.example.com）
- 其余部分指定了web服务器上的某个资源（比如：/image/my-header.jpg）

URI的第二种形式是**统一资源名（URN）**，URN是作为特定资源的唯一名称使用，与当前资源的位置无关。使用与位置无关的RN，就可以将资源位置随意改动。但是URN仍然处于试验阶段，尚未大范围使用。

目前几乎所有的URI都是URL

## 报文

HTTP请求报文和响应报文都包含三部分：起始行、首部字段、主体

**起始行**

报文的第一行就是起始行，在请求报文中主要用来说明做些什么，在响应报文中说明出现了什么情况

**首部字**

起始行后面有零个或者多个首部字段。每个首部字段都包含一个名字和一个值，为了便与解析，两者之间用冒号（:）分割。首部字段以空行结束。

**主体**

空行之后就是可选的主体部分，其中包含了所有类型的数据。请求报文主体中包含要发送给Web服务器的数据，响应主体中装载了要返回给客户端的数据，主体可以包含任意的二进制数据（比如图片，视频，音频等），当然文本是主体中最常见的。

**请求报文**

HTTP请求的起始行有方法，请求地址，和协议版本组成

**例如：**

```
POST /index.php HTTP/1.0
```

HTTP的请求首部字段包含一个名字和一个值，中间又冒号(:)分割

常见的请求首部字段包含User-Agent、HOST、Accept、Accept-language、Cache-Control、Connection、Date、Pragma、Transfer-Encoding、Upgrade、Via等

**例如：**

```
Host:example.com
User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
```

**PHP实例：**

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

HTTP响应起始行有三部分组成，协议HTTP版本、状态码、描述发送状态行

**例如：**

```
HTTP/1.O 404 Not Found
```

HTTP响应首部字段主要是不能放到状态行的附加信息

**例如：**

```
Server: Nginx/1.6.3
Content-length: 403
Content-type: text/html
Date:Mon, 05 Sep 2016 07:05:22 GMT
```

**PHP实例：**

```php
<?php
    header('HTTP/1.0 200 Ok');
    header('Cache-Control: no-cache, must-revalidate');//无缓存
    header('Content-Type: image/png');
    header('Expires: ' . gmdate(DATE_RFC822, time()-3600*24));//过期时间
?>
```

**如何自定义头域和POST数据并获取？**

头域的信息验证在很多的API调用中用到，如下头域中添加了user-appKey和user-appSecret，并POST一组数据

**PHP实例：**

```php
<?php
	$headers = array(
	    'POST /index.php HTTP/1.0',//该行不影响请求，可删除
	    'Content-Type: multipart/form-data;charset="utf-8"',//注意这里的参数设置
	    'User-AppKey: youAppKey',//自定义的头域
	    'User-AppSecret: youAppSercret'
	);

	$post_data = array(
	    'uid' => '10000',
	    'nickName' => '李华',
	);

	/*
	*  如果Content-Type设置为application/x-www-form-urlencoded，需$post_data以urlencode形式
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

使用EGPCS标识获取即：
```
$_COOKIE、$_GET、$_POST、$_FILES、$_SERVER、$_ENV
```

**PHP实例：**

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

**输出结果：**

```
AppKey: youAppKey
AppSecret: youAppSercret
```
nickName: 李华