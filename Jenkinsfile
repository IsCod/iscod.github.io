//
pipeline {
    agent any

    stages {
        stage('pull coding') {
            steps {
            	checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/iscod/iscod.github.io.git']])
                echo 'pull coding success'
            }
        }
        stage('go build') {
            steps {
                echo 'Hello World'
            }
        }
        stage('docker build') {
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