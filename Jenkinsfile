pipeline {
    agent none

    stages {
        stage('Build') {
            agent any
            steps {
                echo 'Run bazel build on gateway service target'
                sh 'export GOPRIVATE=github.com/Melon-Network-Inc/common && bazel build //...'
            }
        }
        stage('Test') {
            agent any
            steps {
                echo 'Run bazel test on gateway service target'
                sh 'export GOPRIVATE=github.com/Melon-Network-Inc/common && bazel test //...'
            }
        }
        stage('Release') {
            agent any
            when { branch "main" }
            steps {
                echo 'Deploying the gateway service application to Production.'
                sh 'screen -S gateway-host  -d -m -c /dev/null -- sh -c "export JENKINS_NODE_COOKIE=dontKillMe; export GOPRIVATE=github.com/Melon-Network-Inc/common; make run; exec sh"'
            }
        }
    }
}