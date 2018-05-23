# HTTP识别认证与安全

## 客户端识别与COOKIE机制

服务器可能同时与数千上百个客户端进行对话, 这些服务器通常需要记录下它们与谁交谈, 而不会认为所有的请求都来自匿名客户端

### 个性化接触

HTTP最初是一个匿名, 无状态的请求响应协议

服务器处理来自客户端的请求, 然后向客户端回送一条响应。web服务器几乎没有什么信息可以判断是那个用户发送的请求, 也无法记录来访用户的请求序列

现代的web站点希望能够提供个性化的接触, 例如:

- 个性化的问候
- 有的放矢的推荐
- 管理信息的存档
- 记录回话

早期的web站点设计者都有自己的用户识别技术, 每种技术都有优劣势

- 承载用户身份信息的HTTP首部
- 客户端IP地址跟踪, 通过用户的IP地址对其进行识别
- 用户登录, 用认证方式来识别用户
- 胖URL, 一种在URL中嵌入识别信息的技术
- cookie, 一种功能强大且高效的持久身份识别技术

### HTTP首部

承载用户相关信息的HTTP首部

 HTTP首部     |  首部类型  | 描述
------------- | -------- | -------------------------------
From          |    请求   | 用户的E-mail地址
User-Agent    |    请求   | 用户的浏览器软件
Referer       |    请求   | 用户是在这个页面上依照连接跳转过来的
Authorization |    请求   | 用户是在这个页面上依照连接跳转过来的
Client-ip     | 请求(扩展) | 客户端的IP地址
Cookie        | 请求(扩展) | 服务器产生的ID标签

From 包含了用户的E-mail地址, 每个用户都有不同的E-mail地址, 所有在理想情况下可以将这个地址作为可行的源端来识别用户. 但由于担心不道德的服务器会收集这些E-mail地址, 用于垃圾邮件散发所以很少有浏览器会发送From首部

User-Agent 首部可以将用户的浏览器的先关信息告知服务器, 包括程序的名称和版本, 通常还包含操作系统的相关信息. 要实现特定浏览器及其属性的良好互操作时, 这个首部非常有用, 但是它没有为识别特定用户提供帮助

Referer 首部提供了用户来源页面的URL, 它自身并不能完全标识用户

### 客户端IP

早期的web先锋曾尝试客户端IP作为一种标识形式使用, 如果每个用户都有不同的IP地址, IP地址也很少发生变化, 而且服务器可以判断出每条请求的客户端IP地址

但是, 使用客户端IP有很多缺点, 限制将其作为用户识别技术的效能

- 客户端IP描述的是所用的机器而不是用户, 如果多个用户使用一台机器, 那么就无法对其区分了
- 很多网络提供商都会在用户登录时为其分配动态IP地址, 用户每次登录都会得到以一个不同的地址, 因此web服务器不能假设IP地址在各登录会话之间识别用户
- 为了提高安全, 并对稀缺得地址资源进行管理, 很多用户是通过网络地址转换 (NAT) 防火墙来浏览内容, 这些NAT防火墙隐藏了防火墙后哪些实际客户端IP地址, 将实际客户端IP地址转换成了一个共享的防火墙IP (和不同的端口)
- HTTP代理和网关通常会打开一个新的, 到原始服务器的TCP连接, web服务器将看到的是代理服务器的IP而不是客户端的


### 用户登录



## 基本认证机制

## 摘要认证

## 安全的HTTP

基本认证, 摘要认证和报文完整性检查对于一般的网路事务来说这些方法很好用, 但是对于大规模的购物银行事务, 或者对访问机密数据来说, 并不够强大

这些更重要的事务需要将HTTP和数字加密技术结合使用才能确保完全

HTTP的安全版本同时需要高效可移植且易于管理并且提供下列功能

- 服务器认证 (客户端知道它们是与真正的而不是伪造的服务器通话)
- 客户端认证 (服务器知道它们是与真正的而不是伪造的客户端通话)
- 完整性 (客户端与服务器的数据不会被修改)
- 加密 (客户端与服务器的对话是私密的, 无需担心被窃听)
- 效率 (一个运行足够快的算法, 以便低端的客户端和服务器使用)
- 普适性 (基本上所有的客户端和服务器都支持这种协议)
- 管理的可扩展性 (在任何地方的任何人都可以立即建立安全通信)
- 适应性 (能够支持当前最知名的安全方法)
- 在社会的可行性 (满足社会的政治文化需求)

### HTTPS

HTTPS是目前最流行的HTTP安全形式, 由网景公司首创, 所有主要的浏览器和服务器都支持该协议

HTTPS 方案的URL开头是https://而不是http://开头, 可以根据方案确定使用的是HTTPS还是HTTP

使用HTTPS时, 所有的HTTP请求和响应数据在发送到网络前, 都进行了加密

HTTPS在HTTP协议下面提供一个传输级的密码安全层, 一般使用SSL或者TLS

由于大部分困难的编码和解码工作由SSL库完成, 所以web客户端和服务器在使用安全的HTTP时无需做太多的修改协议处理逻辑

大多数情况下, 只需要用SSL的输入/输出调用取代TCP的调用, 再增加几个调用来配置管理安全信息就可以了

### 数字加密

介绍数字加密前，你要先了解以下几个内容

### 密码

对文本进行编码，使偷窥者无法识别的算法

### 密钥

改变密码行为的数字化参数

### 对称密钥加密系统

编/解码使用相同的密钥的算法

### 不对称密钥加密系统

编/解码使用不同的密钥的算法

### 公开密钥加密系统

一种能够使数百万计算机便捷地发送机密报文的系统

### 数字签名

用来验证报文未被伪造或篡改的校验和

### 数字证书

由一个可信的组织验证和签发的识别信息

密码学是对报文进行编码/解码的机制和技巧。密码学除了加密报文，还可以验证某个报文或某个事物确实出自你手

## 密码学

密码学基于一种名为密码的密码代码。密码是一套编码方案—一种特殊的编码方式和一种相应解码方式的结合体。加密以前的原始报文称为明文，使用了密码之后的编码报文通常被称为密文。随着密码机的出现，这些机器可以用复杂的多的密码来快速，精确的对报文进行编解码。

用密码生成保密信息已经有千年历史，传说凯撒曾使用一种三字符循环位移密码，报文中的每个字符都由字母表中三个位置之后的字符来取代，这种位移密码演变出多种方法比如维吉尼亚密码等。

**凯撒密码示例：**

```php
<?php
  $keys = array('A' => 'N','B' => 'O','C' => 'P','D' => 'Q','E' => 'R','F' => 'S','G' => 'T','H' => 'U','I' => 'V','J' => 'W','K' => 'X','L' => 'Y','M' => 'Z','N' => 'A','O' => 'B','P' => 'C','Q' => 'D','R' => 'E','S' => 'F','T' => 'G','U' => 'H','V' => 'I','W' => 'J','X' => 'K','Y' => 'L','Z' => 'M','a' => 'n','b' => 'o','c' => 'p','d' => 'q','e' => 'r','f' => 's','g' => 't','h' => 'u','i' => 'v','j' => 'w','k' => 'x','l' => 'y','m' => 'z','n' => 'a','o' => 'b','p' => 'c','q' => 'd','r' => 'e','s' => 'f','t' => 'g','u' => 'h','v' => 'i','w' => 'j','x' => 'k','y' => 'l','z' => 'm');

  $data = 'example';

  $data_arr = str_split(trim($data));

  $new_data = '';

  if (is_array($data_arr)) {
    foreach ($data_arr as $value) {
      if (array_key_exists($value, $keys)) {
        $new_data .= $keys[$value];
      }else{
        $new_data .= $value;
      }
    }
  }

  reutrn $new_data;
?>
```

### 密码机

最初，人们需要自己进行编码和解码，所以起初密码是相当简单的算法。因为密码很简单，所以人们通过纸笔和密码书就可以进行编解码了，但聪明人也可以相当容 易地“破解”这些密码。

随着技术的进步，人们开始制造一些机器，这些机器可以用复杂得多的密码来快速、 精确地对报文进行编解码。这些密码机不仅能做一些简单的旋转，它们还可以替换字符、改变字符顺序，将报文切片切块，使代码的破解更加困难。(最著名的机械编码机可能就是第二次世界大战期间德国的 Enigma 编码机了。尽管 Enigma 密码非常复杂，但阿兰 · 图灵(Alan Turing)和他的同事们在 20 世纪 40 年代初期就可以用最早的数字计算机 破解 Enigma 代码了)

编码算法和编码机有可能落入敌人的受众，所以大部分机器上都有一些号盘。可以将其设置为大量不同的值以改变密码的工作方式。即使机器被盗，没有正确的号盘设置（密钥值），解码器也无法工作。设置这些密码的参数被称为密钥(key)。要在密机中输入正确的密钥，解密过程才能正确进行。

### 数字密码

随着数字计算的出现，出现了以下两个主要的进展。

从机械设备的速度和功能限制中解放出来，使复杂的编 / 解码算法成为可能。

支持超大密钥成为可能，这样就可以从一个加密算法中产生出数万亿的虚拟加密算法，由不同的密钥值来区分不同的算法。密钥越长，编码组合就越多，通过随机猜测密钥来破解代码就越困难。

与金属钥匙或机械设备中的号盘设置相比，数字密钥只是一些数字。这些数字密钥 值是编 /解码算法的输入。编码算法就是一些函数，这些函数会读取一块数据，并 根据算法和密钥值对其进行编 / 解码。

### 对称密钥加密技术

很多数字加密算法被称为对称密钥加密技术，这是因为他们在编码时使用的密钥和解码时的一样。流行的对称密钥加密算法包括：DES、Triple-DES、RC2和RC4

对称加密技术的缺点之一就是发送者和接收者在相互对话之前，一定要有一个共享的保密密钥。如果每对通信都需要自己使用私有的密钥，如果有N个节点，每个节点都要和其他所有N-1个节点进行安全对话，总共需要大约N2个保密的密钥。

### 公开密钥加密技术

公开密钥加密技术没有为每对主机使用单独的加密/解密密钥，而是使用了两个非对称密钥，一个用来对主机报文编码，另一个用来对主机报文进行解码。编码密钥是众所周知的（也就是公开密钥加密这个名字的由来），但只有主机才知道私有的解密密钥。这样，每个人都能找到某个特定主机的公开密钥，密钥的建立变得更加简单。但是解码的密钥是保密的，因此只有接收端才能对发送给它的报文进行解码

RSA
所有公开密钥非对称加密系统所面临的共同挑战是，要确保即使有人拥有了下面所有的线索，也无法计算出保密的私有密钥：

### 公开密钥（是公开的，所有人都可以获取）

- 一小片拦截下的密文（可通过对网络的嗅探获取）
- 一条报文与之相关的密文（对任意一段文本运行加密器就可以得到）

RSA算法就是满足了所有这些条件的流行的公开密钥加密系统，它是在MIT发明后由RSA数据安全公司将其商业化。即使有了上述所有条件，RSA算法自身，甚至RSA实现的源代码，破解代码找到相应的私有密钥的难度仍相当于对一个极大的数进行质因子分解的困难程度，这被认为是所有计算机科学中最难的问题之一。因此如果你发现一种快速将一个极大的数分解为质因数的方法，不仅能够入侵瑞士银行账户系统，而且可以获得图灵奖了

RSA加解密示例：

```php
<?php
  class Rsa_encrypt{
    private static $private_key = '-----BEGIN RSA PRIVATE KEY-----
      MIICXAIBAAKBgQDCKunY6xL9TyNmynkRp5qLG/szrIvyfTku6ZN2TcXhIuVAeOem
      oYSXX1S3thPaW/JEIWR5RrSUDhULA0d57s7sRCYlopNe8Blenvz81QU2ttFmc070
      fHhPbspWtPizbo3xYW74Y/0vQOpiOh5qNz6c7dSkzd7GKx7Z+hm9FJQK/QIDAQAB
      AoGBAKWAGIUJkc0SCHXUPS/cMXFDL3HTMBJHxFcFRuj+z5zftpKmu6UfZTn1Suuw
      Kenkl3KVF+P7bW4JNsyFRgZblEj4iT84XJCyTt8ztdFUomTwQu4/d8MVt1euI2S4
      NLw0m2AJHPMDTvzwJlys1XPsAYsDp9GHmP5rm4B3gO4zIoQNAkEA4reLmyiIyR/e
      NhmlxMP6an8ZGEM881g3FSZeiujQDUdUDX04wq5x7Kl8hYebhau6WGv7kQENlNfP
      dXbzhQ2JdwJBANs/HaPFE4VqGBtlt1GR+QYuqfjnrNwvzZtk7KZJOfKUg8NGQmwR
      szgEeUxdyzoZEvaf5r8Jygkb8to8e3EerCsCQFJ9H8FrZSFwg+RBPqwx9hnrdpD6
      XeHYVepPFJUMEi7SpgVma1GCMRc/r3vSFEb1bY6gc16V+IAQaX4+smnVvA8CQGNe
      NMnP/Wv/TNPGAxL2TN5Pcfv8zKyzAcYHNPacw6W9SAbOJjaiww6FgJBrBjvbt2uN
      x2AYSLheMXBV70CyvScCQDvjgNgc/kkqkfW9r3n6qtfWAw/rX+fMFF6t7fa8h6kr
      fMNyH+ULbWCjiSdabuQuAKBDP9rGmdR/62mYp39l8bE=
      -----END RSA PRIVATE KEY-----';//私钥

    private static $public_key = '-----BEGIN PUBLIC KEY-----
      MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCKunY6xL9TyNmynkRp5qLG/sz
      rIvyfTku6ZN2TcXhIuVAeOemoYSXX1S3thPaW/JEIWR5RrSUDhULA0d57s7sRCYl
      opNe8Blenvz81QU2ttFmc070fHhPbspWtPizbo3xYW74Y/0vQOpiOh5qNz6c7dSk
      zd7GKx7Z+hm9FJQK/QIDAQAB
      -----END PUBLIC KEY-----';//公钥

    public function __construct() {
      //construct
    }

    /**
    * 使用密钥编码
    */
    public function encode($data = array(), $block_size = 200) {
        $encrypted = '';

        $data = (string)json_encode($data);
        $plainData = str_split($data, $block_size);
        if (!$plainData) return FALSE;

        $privateKey = openssl_pkey_get_private($this->private_key);

        if (!$privateKey) return FALSE;

        foreach($plainData as $chunk)
        {
            $partialEncrypted = '';

            //using for example OPENSSL_PKCS1_PADDING as padding
            $encryptionOk = openssl_private_encrypt($chunk, $partialEncrypted, $privateKey, OPENSSL_PKCS1_PADDING);

            if($encryptionOk === false){return FALSE;}//also you can return and error. If too big this will be false
            $encrypted .= $partialEncrypted;
        }

        return base64_encode($encrypted);
    }

    /**
    *使用公钥解码
    *解码来自客户端公钥的密文
    * @param string $str密文
    * @return $data 明文变量
    public function decode($sign = '', $block_size = 256){
        $decrypted = '';

        //decode must be done before spliting for getting the binary String
        $data = str_split(base64_decode($sign), $block_size);
        if (!$data) return $decrypted;

        $publicKey = openssl_pkey_get_public($this->public_key);
        if(!$publicKey) return $decrypted;

        foreach($data as $chunk)
        {
            $partial = '';
            //be sure to match padding
            $decryptionOK = openssl_public_decrypt($chunk, $partial, $publicKey, OPENSSL_PKCS1_PADDING);
            if($decryptionOK === false){return FALSE;}//here also processed errors in decryption. If too big this will be false
            $decrypted .= $partial;
        }

        return json_decode($decrypted);
    }
  }
?>
```
