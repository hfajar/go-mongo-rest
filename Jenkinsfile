Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                echo 'masuk sini'
                sh 'go version'
		        sh 'go run main.go'
            }
        }
    }
}
