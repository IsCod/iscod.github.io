# kubernetes

`kubernetes`是一个用于容器编排的引擎，同时提供容器化应用的自动化部署、 扩缩和管理

## resources

```
# 查看资源名称与对应的 apiVersion
kubectl api-resources
```

### Pod

```bash
# 运行一个busybox的pod
kubectl run -i -t busybox --image=busybox
```

#### Probe健康检查

容器可配置存活（Liveness）、就绪（Readiness）和启动（Startup）三种探针形式

* 存活（Liveness）: 用来确定什么时候重启容器，例如，存活探针可以监测到应用死锁情况，重启这类容器可以提高应用的可用性
* 就绪（Readiness）: 用来确定容器何时可以接收请求流量，提供服务
* 启动（Startup）: 可以用来保护慢启动容器，

> 存活探针是一种从应用故障中恢复的强劲方式，但应谨慎使用。 你必须仔细配置存活探针，确保它能真正标示出不可恢复的应用故障，例如死锁。

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    name: probe-test 
  name: probe-test
spec:
  containers:
  - name: nginx
    image: nginx
    livenessProbe: # 定义一个存活探针，http请求正确返回才可用，一般根据业务逻辑确定
      httpGet:
        path: /
        port: 80
      failureThreshold: 30 # 探针连续失败了 n 次 认为失败
      initialDelaySeconds: 15 # 容器启动后要等待多少秒后才启动启动、存活和就绪探针
      periodSeconds: 3 # 指定 每隔 3 秒执行一次存活探测
    readinessProbe: # 定义一个就绪探针，http请求正确返回才可用
      tcpSocket:
        port: 80
      initialDelaySeconds: 5
      periodSeconds: 10
    startupProbe: # 启动探针
      httpGet:
        path: /
        port: 80
      failureThreshold: 30 # 探针连续失败了 n 次 认为失败
      periodSeconds: 10
# 有启动探测，应用将会有最多 5 分钟（30 * 10 = 300s）的时间来完成其启动过程。
# 一旦启动探测成功一次，存活探测任务就会接管对容器的探测，对容器死锁作出快速响应。
# 如果启动探测一直没有成功，容器会在 300 秒后被杀死，并且根据 restartPolicy 来执行进一步处置。
```

### deployment

```bash

# 根据一个yaml文件创建一个 deployment 资源
kubectl create deployment user-uservice-deployment -o user-uservice-deployment.yaml

# apply根据一个yaml文件更新一个 deployment 资源
kubectl apply -f user-service-deployment.yaml

# deployment 更新镜像，常与ci/cd结合更新程序
kubectl set image deployment/deployment-name nginx=nginx:1.24

# 查看 deployment 发布历史版本
kubectl rollout history deployment user-service-deployment
# 查看 deployment 发布历史版本详情
kubectl rollout history deployment user-service-deployment --revision=4

### deployment 回滚, 注意多次回滚只是在最后两个版本相互切换
kubectl rollout undo deployment user-service-deployment

### deployment 回滚到指定版本
kubectl rollout undo deployment user-service-deployment --to-revision=4
```

### statefulSet

有状态应用管理，使用于有状态的应用场景，如：redis-cluster, mongoDB集群、kafka集群、Eureka等

创建 `statefulSet` 需存在对应的无头`service`

```yaml
## redis for service
apiVersion: v1
kind: Service
metadata:
    name: redis-service
    labels:
        name: redis-service
    namespace: default
spec:
    type: ClusterIP
    clusterIP: None # 标识为无头服务，无头服务即不会生成clusterIP
    selector:
        name: redis-pod
    ports:
      - name: "redis-port"
      port: 6379
      protocol: TCP
      targetPort: 6379
```

```yaml
## redis for statefulSet
apiVersion: apps/v1
kind: StatefulSet
metadata:
    name: redis-sts
    namespace: default
    labels:
        name: redis-sts
spec:
    replicas: 2
    serviceName: redis-service # 这里是上一步创建的service名称
    selector:
        matchLabels:
          name: redis-pod # 必须匹配到template的labels
    template:
        metadata:
            labels:
              app: redis-app
              name: redis-pod
        spec:
            containers:
            - image: redis:latest
              imagePullPolicy: IfNotPresent
              name: redis-container
              ports:
              - containerPort: 6379
                name: redis-port
    updateStrategy:
        type: RollingUpdate
        rollingUpdate:
            partition: 0 #指定灰度发布
```

### DaemonSet

在匹配的节点或者所有节点创建一个副本，适用于如：节点日志收集、CNI、calico、ingress、node exporter 、flannel

```
# 给节点添加 label: ingress-nginx=true
kubectl label node k8s-node000012 ingress-nginx=true
```

```yaml
## DaemonSet nginx
apiVersion: apps/v1
kind: DaemonSet
metadata:
    name: ingress-nginx
    labels:
        name: ingress-nginx-ds
spec:
    selector:
        matchLabels:
            name: ingress-node
    template:
        metadata:
            labels:
                name: ingress-node
        spec:
            nodeSelector:
                ingress-nginx: "true" # 根据label，选择node, 不设置则表示选择所有节点
            containers:
                - image: nginx:latest
                  imagePullPolicy: IfNotPresent
                  name: ingress-nginx-container
```

### service

service类型：`ExternalName`, `LoadBalancer`, `NodePort`, `ClusterIP`, `ClusterIP-none`

### label & selector

#### label

`label`可以给任意资源(service,pod,node等)添加一个标签分组

```bash
kubectl label pod busybox version=v1 # 添加一个label
kubectl label pod busybox version=v2 --overwrite # 修改一个label
kubectl get pod -l version=v2 # 筛选version=v2的pod
kubectl get pod -l 'version in (v2, v1)' # 通过in筛选version=v1和version=v2的资源
kubectl label pod busybox version- # 删除一个label
```

#### selector

`selector`主要用于筛选符合标签的资源。
例如: 通过`service.spec.selector`选择对应的`pod`, `Deployment.spec.template.spec.nodeSelector`选择在合适的`node`创建pod

```yaml
apiVersion: v1
kind: Service
metadata:
    name: user-service
    labels:
        name: user-service
spec:
    selector:
        name: user-service # 选择name=user-service的pod
```

### ConfigMap && Secret

#### ConfigMap

```bash
kubectl create configmap myconf --from-file conf.yml # 通过配置文件创建 configmap
```

如何`deployment`中使用configmap？

```yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx
  namespace: default
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx
        volumeMounts:
        - mountPath: /var/conf # 容器中的保存路径
          name: config # 与volumes中的name一直
          readOnly: true
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: conf.yml
            path: conf.yml
          name: myconf # configmap中的名字
        name: config

```

#### Secret

`secret` 主要用于保存少量敏感信息例如密码、令牌或密钥的对象。 这样的信息可能会被放在 Pod 规约中或者镜像中。

Secret 类似于 ConfigMap 但专门用于保存机密数据。

```
# secret 提供直接创建 docker-registry 认证
kubectl create secret docker-registry myregister --docker-server=DOCKER_REGISTRY_SERVER --docker-username=xxxx --docker-password=xxxx --docker-email=xxxx

# 创建 nginx 的 tls 证书
kubectl create secret tls NAME --cert=path/to/cert/file --key=path/to/key/file
```



### scale && autoscale

``` bash
# 设置一个deployment的autocale
kubectl autoscale deployment user-service-deployment --min=3 --max=6

# 查看当前自动调度列表
kubectl get hpa -o wide
```

## ingress

[ingress-nginx](https://kubernetes.github.io/ingress-nginx)和[nginx-ingress](https://github.com/nginxinc/kubernetes-ingress)是两个不同的项目。

`ingress-nginx`是k8s官方维护的ingress，
`nginx-ingress`是nginx官方维护的ingress，两者配置基本相同，在k8s中推荐`ingress-nginx`

```bash
helm pull ingress-nginx/ingress-nginx # 下载ingress-nginx
tar xf ingress-nginx-4.7.1.tgz # 解压
vim ingress-nginx/values.yaml # 修改模版配置文件
```

下载`ingress-nginx`后我们需要根据自己的业务需要，修改模版(values.yaml)文件的几个关键位置

1. controller,admissionWebhooks的image地址，选择阿里云等镜像
2. `hostNetwork`修改为`true`，推荐使用
3. `dnsPolicy`选择`ClusterFirstWithHostNet`模式
4. `nodeSelector`添加标识如`ingress-nginx: true`，这样就能指定节点安装
5. `kind`指定为`DaemonSet`类型

```bash
# 启动ingress-nginx
helm install ingress-nginx -n ingress-nginx . # 在ingress-nginx目录下执行

# 创建一个简单的ingress
kubectl create ingress ingerss-nginx-svc --class=nginx --rule="ingerss.test.com/*=nginx-svc:80"
```

```yaml
# ingress的config yaml文件
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
    name: ingerss-nginx-svc
    namespace: default # service的namespace相同
spec:
    ingressClassName: nginx # 查看已创建的ingressClassName: kubectl get ingressclasses.networking.k8s.io
    rules:
        - host: ingerss.test.com
    http:
        paths:
        - backend:
            service:
                name: nginx-svc # 代理的服务名
                port:
                    number: 80
                path: /
                pathType: Prefix
```

```bash
curl ingerss.test.com # 域名解析后或添加hosts后，测试域名
```

> 访问ingress代理的域名，域名需解析到安装`ingress-nginx-controller` pod 的`node`节点才可访问。详见：[nodeSelector](https://iscod.github.io/#/devops/kubernetes?id=ingress)配置

#### annotations

ingress-nginx通过配置`annotations`可以实现常用的服务功能：
[速率限制](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#rate-limiting)
[重定向](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#permanent-redirect)
[白名单](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#whitelist-source-range)
[代理重定向](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#proxy-redirect)

```yaml
# 常用 metadata.annotations配置
nginx.ingress.kubernetes.io/rewrite-target: /$2 # 转发
nginx.ingress.kubernetes.io/use-regex: "true" # 配置转发时需要的正则表达式，path配置正在规则
nginx.ingress.kubernetes.io/limit-rps: 1 # 速率限制
nginx.ingress.kubernetes.io/permanent-redirect: https://www.google.com #重定向
nginx.ingress.kubernetes.io/temporal-redirect: https://www.google.com # 临时重定向
nginx.ingress.kubernetes.io/whitelist-source-range: 10.0.0.0/24,172.10.0.1 # 白名单，多个ip `,` 分割
```

```yaml
# rewrite配置
# foo.com/api -> foo.com
# foo.com/api/new -> foo.com/new
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
    annotations:
        nginx.ingress.kubernetes.io/rewrite-target: /$2
        nginx.ingress.kubernetes.io/use-regex: "true"
    generation: 3
    name: rewrite
    namespace: default
spec:
    ingressClassName: nginx
    rules:
    - host: foo.com
        http:
        paths:
        - backend:
          service:
            name: nginx-svc
            port:
                number: 80
          path: /api(/|$)(.*)
          pathType: Prefix
```

#### ingress-nginx实现原理？

ingress-nginx会将用户的配置，注入到启动的节点`pod`中的nginx配置文件，从而实现代理，我们看下pod中的配置文件

```bash
# 查看ingress-nginx pod中的nginx.conf配置文件， 注意pod名称
kubectl exec -it ingress-nginx-controller-9qwnw --namespace ingress-nginx cat /etc/nginx/nginx.conf |grep foo.com -A 10
```

## metrics-server

## [prometheus](https://github.com/prometheus/prometheus)

#### install

```bash
git clone https://github.com/prometheus-operator/kube-prometheus.git #这里选择kube-prometheus版本
cd kube-prometheus/manifests
kubectl apply --server-side -f setup
kubectl wait --for condition=Established --all CustomResourceDefinition --namespace=monitoring
kubectl apply -f .
```

[kube-prometheus](https://github.com/prometheus-operator/kube-prometheus)

#### 数据可视化

安装完成后，prometheus为我们创建了一系列的服务其中包括：

`alertmanager-main` 是我们的需要的`alert`告警配置
`grafana` 是`prometheus`数据的可视化 webUI
`prometheus-k8s`  是相关的`prometheus`自身配置，如 alert、Rules、Configuration

通过`kubectl get svc --namespace monitoring`可以查看生产的服务

# prometheus生成的相关服务:

    # prometheus生成的相关服务：
    alertmanager-main       ClusterIP   10.108.145.122   <none>        9093/TCP,8080/TCP            19h
    grafana                 ClusterIP   10.98.163.188    <none>        3000/TCP                     19h
    prometheus-k8s          ClusterIP   10.105.23.112    <none>        9090/TCP,8080/TCP            19h


而我们想通过域名访问，需要通过[ingress](https://iscod.github.io/#/devops/kubernetes?id=ingress)进行服务的反向代理

```bash
# 创建ingress, namespace必须和服务在同一命名空间, 留意alertmanager-main,prometheus-k8s的端口号
kubectl create ingress prometheus --class=nginx --rule="grafana.test.com/*=grafana:3000" --rule="alert.test.com/*=alertmanager-main:9093" --rule="prometheus.test.com/*=prometheus-k8s:9090" --namespace=monitoring
```

#### 告警配置

prometheus提供多种告警通知包括常用的：
[email_config](https://prometheus.io/docs/alerting/latest/configuration/#email_config) 、
[wechat_config](https://prometheus.io/docs/alerting/latest/configuration/#wechat_config)


*邮件配置*

在`alertmanager-secret.yaml`添加我们的邮件信息

```yaml
"global": #global中添加邮件服务器信息
  "resolve_timeout": "6m"
  "smtp_from": "xxxx@163.com"
  "smtp_smarthost": "smtp.163.com:465" # smtp根据邮箱服务商配置
  "smtp_hello": "prometheus"
  "smtp_auth_username": "xxxx@163.com"
  "smtp_auth_password": "xxxx"
  "smtp_require_tls": false #注意，163不支持tls必须添加为false才行
"receivers":
    - "name": "Default"
      "email_configs": #receiversz中添加接受者email_configs配置
      - "to": "xxxx@163.com"
        "send_resolved": true
```

*微信配置*

#### 架构图

![prometheus](https://iscod.github.io/images/prometheus.png)

## istio

### 服务网格