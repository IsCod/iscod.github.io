//
pipeline {
    agent any

    tools { go '1.19' }

    stages {
        stage('pull coding') {
            steps {
            	checkout scmGit(branches: [[name: '${tag}']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/iscod/iscod.github.io.git']])
                echo 'pull coding success'
            }
        }
        stage('Build') {
          steps {
            // Output will be something like "go version go1.19 darwin/arm64"
            sh 'go version'
          }
        }
        stage('DockerBuild') {
            steps {
                echo 'docker build image success'
            }
        }
        stage('image push') {
            steps {
                echo 'docker image push registry'
            }
        }
        stage('k8s update') {
            steps {
                echo 'k8s update success'
            }
        }
    }
}