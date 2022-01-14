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
                cd bin ./setup
            }
        }
        stage("deploy"){
            steps {
            echo 'in deploy stage'
            }
        }
    }
}