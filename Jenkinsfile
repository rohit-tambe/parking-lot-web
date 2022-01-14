pipeline {
    agent any

    stages{
        stage("build"){
            steps {
                echo 'in build stage'
                
            }
        }
        stage("test"){
            steps {
                echo 'in test stage'
                sh cd /bin ./setup
                echo 'in test stage complete'
            }
        }
        stage("deploy"){
            steps {
            echo 'in deploy stage'
            }
        }
    }
}