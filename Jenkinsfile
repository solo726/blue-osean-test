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
    triggers {
          pollSCM('H/3 * * * *')
      }
  stages {
    stage('prepare') {
      parallel {
        stage('debug info') {
          steps {
            sh '''
hostname

docker version

go version 

kubectl version

kubectl get pods

helm version

ip a

uname -a

whoami

pwd

set


mkdir /pipeline-info

echo `git rev-parse --short HEAD` > /pipeline-info/git-commit


echo "git-commit:`cat /pipeline-info/git-commit`"


'''
          }
        }
        stage('code fmt') {
          steps {
            sh '''cp -r $WORKSPACE /tmp/blue-osean-test

ls -la /tmp/blue-osean-test

cd /tmp/blue-osean-test

go fmt  '''
          }
        }
        stage('image build') {
          steps {
            sh '''
gitCommitVersion=`cat /pipeline-info/git-commit`

echo "gitCommitVersion:${gitCommitVersion}"

imageName="solo726/blue-osean-test:${gitCommitVersion}"

echo "imageName:${imageName}"


echo "${imageName}" > /pipeline-info/image-name

cd /tmp/blue-osean-test

docker build -f deploy/docker/Dockerfile . --tag ${imageName}

'''
          }
        }
      }
    }
    stage('unit test') {
      steps {
        sh '''
imageName=`cat /pipeline-info/image-name`

echo "${imageName}"


docker run --rm ${imageName} go test -v -cover=true /go/src/blue-osean-test/main_test.go /go/src/blue-osean-test/main.go

'''
      }
    }
    stage('image push') {
      steps {
        withCredentials(bindings: [usernamePassword( credentialsId: 'docker-login', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
          sh '''
            docker login -u ${USERNAME} -p ${PASSWORD}
            docker push `cat /pipeline-info/image-name`
          '''
        }

      }
    }
    stage('deploy') {
      steps {
        sh '''cd /tmp/blue-osean-test

ls

helm init

helm upgrade --install --namespace blue-osean-test --debug --set images=`cat /pipeline-info/image-name` blue-osean-test deploy/helm/blue-osean-test'''
      }
    }
  }
}