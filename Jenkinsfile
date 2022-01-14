pipeline {
    agent any

    stages{
        
        stage("test"){
            steps {
                echo 'in test stage'
                sh 'cd /bin ./setup'
                echo 'in test stage complete'
            }
        }
        stage("build"){
            steps {
                 sh 'cd /bin ./build'
                echo 'in test stage complete'
            }
        }
    }
}