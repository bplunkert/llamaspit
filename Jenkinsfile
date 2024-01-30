pipeline {
  agent any

  environment {
    OLLAMA_HOST = 'http://inference:11434'
  }

  stages {
    stage('Build') {
      steps {
        sh 'docker compose build'
      }
    }

    stage('Test') {
      steps {
        sh 'docker compose run llamaspit go test'
      }
    }
  }
  post {
    always {
      sh 'docker compose down'
    }
  }
}
