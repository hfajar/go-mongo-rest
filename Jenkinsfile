Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh 'go version'
		        sh 'go run main.go'
            }
        }
    }
}
