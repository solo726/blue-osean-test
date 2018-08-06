pipeline {
  agent any
  stages {
    stage('prepare') {
      parallel {
        stage('prepare') {
          steps {
            echo 'debug info'
          }
        }
        stage('code fmt') {
          steps {
            echo 'code fmt'
          }
        }
        stage('image build') {
          steps {
            echo 'image build'
          }
        }
      }
    }
    stage('unit test') {
      steps {
        echo 'unit test'
      }
    }
    stage('image push') {
      steps {
        echo 'image push'
      }
    }
    stage('deploy') {
      steps {
        sh 'echo deploy'
      }
    }
  }
}