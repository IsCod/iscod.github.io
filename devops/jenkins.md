# Jenkins

## 流水线(pipeline)

Jenkins 流水线 (或简单的带有大写"P"的"Pipeline") 是一套插件，它支持实现和集成 continuous delivery pipelines 到Jenkins。

Jenkins 流水线的定义可以到一个文件中 (Jenkinsfile)保存到源代码的版本控制库中

## 代理(agent)

agent 指定了整个流水线或特定的部分, 将会在Jenkins环境中执行的位置，这取决于 agent 区域的位置。该部分必须在 pipeline 块的顶层被定义, 但是 stage 级别的使用是可选的。

#### 参数
`none`

如果在 pipeline 块的顶部设置该参数(既没有全局代理)，该参数将会被分配到整个流水线的运行中，则每个 stage 部分`都需要包含他自己的 agent` 部分。比如: agent any

`any`

在任何可用的代理上执行流水线或阶段。例如: agent any

`docker`

使用给定的容器内执行流水线或阶段

```
stage('Build') {
    agent {
        docker {
            image 'golang'
        }
    }
    steps {
        sh 'go version'
    }
}
```

`dockerfile`

使用从源代码库包含的 Dockerfile 构建的容器

```
agent {
    // Equivalent to "docker build -f Dockerfile.build --build-arg version=1.0.2 ./build/
    dockerfile {
        filename 'Dockerfile'
        dir 'build'
        additionalBuildArgs  '--build-arg version=1.0.2'
    }
}
```

## stages

包含一系列一个或多个 stage 指令, stages 部分是流水线描述的大部分"work" 的位置。 建议 stages 至少包含一个 stage 指令用于连续交付过程的每个离散部分,比如构建, 测试, 和部署。

## steps

steps 部分在给定的 stage 指令中执行的定义了一系列的一个或多个steps。

## 示例
```jenkins
pipeline {
  agent none
  stages {
    stage('Prepare') {
      steps {
        echo 'Prepare...'
      }
    }
    stage('Build') {
      agent {
        docker {
          image 'golang'
        }
      }
      steps {
        checkout([$class: 'GitSCM', branches: [[name: '$COMMIT']], extensions: [], userRemoteConfigs: [[credentialsId: 'ca9968d1-3d82-453c-8380-27a1bb300a2d', url: 'git@github.com:iscod/IsCod.github.io.git']]])
        sh 'ls -lh'
        sh 'go mod download'
        sh 'go build -a -v -o tao-gin'
      }
    }
    stage('Test') {
        agent any
        steps {
            sh 'ls -lh'
        }
    }
    stage('Deploy') {
        agent any
        steps {
            sh 'if ([ "$ENV" != "" ]); then mv conf/gin.$ENV.conf conf/gin.conf; fi'
            sh 'docker build  -t $JOB_BASE_NAME:$BUILD_NUMBER  .'
        }
    }
  }
}
``` 

* 参考
    * [Jenkins](https://www.jenkins.io/zh/doc/book/pipeline/)