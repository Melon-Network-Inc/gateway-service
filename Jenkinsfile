pipeline {
    agent none

    stages {
        stage('Build') {
            agent any
            when { branch "main" }
            steps {
                echo 'Run bazel build on gateway service target'
                sh 'export GOPRIVATE=github.com/Melon-Network-Inc/common && bazel build //...'
            }
        }
        stage('Cleanup') {
            agent any
            when { branch "main" }
            steps {
                echo 'New release is approved. Clean up previous release.'
                sh 'screen -XS gateway-host quit'
            }
        }
        stage('Release') {
            agent any
            when { branch "main" }
            steps {
                echo 'Deploying the gateway service application to Production.'
                sh 'export JENKINS_NODE_COOKIE=dontKillMe; screen -S gateway-host  -d -m -c /dev/null -- sh -c "export GOPRIVATE=github.com/Melon-Network-Inc/common; make run; exec sh"'
            }
        }
    }
}