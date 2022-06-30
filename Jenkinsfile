pipeline {
    agent any
    stages {
        stage('Cleanup') {
            steps {
//                 sh 'go version'
                // TODO 调整顺序，测试通过了再删除现有容器
                sh 'docker stop goblog || true'
                sh 'docker rmi goblog || true'
            }
        }
        stage('Build') {
            steps {
                sh 'docker build -t goblog .'
            }
        }
        stage('Test') {
            steps {
                sh 'docker run --rm -v /var/jenkins_home/workspace/goblog/report:/app/report goblog'
            }
        }
    }
}