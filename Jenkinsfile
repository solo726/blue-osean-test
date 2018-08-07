pipeline {
  agent {
    kubernetes {
      label "jenkins-job-jnlp-${UUID.randomUUID().toString()}"
      yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: jnlp
    image: \'solo726/jenkins:jnlp-slave\'
    tty: true
    volumeMounts:
      - name: docker-sock
        mountPath: /var/run/docker.sock
      - name: kube-config
        mountPath: /root/.kube
  volumes:
    - name: docker-sock
      hostPath:
        path: /var/run/docker.sock
    - name: kube-config
      secret:
        secretName: kubeconfig
        items:
          - key: config
            path: config
'''
    }

  }
  stages {
    stage('prepare') {
      parallel {
        stage('debug info') {
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
