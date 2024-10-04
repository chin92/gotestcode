

// Function to compare version strings
boolean isVersionLessThan(String v1, String v2) {

    def credentialsId = 'your-credentials-id'

    // Using the SSH key for Git operations
    withCredentials([sshUserPrivateKey(credentialsId: credentialsId, keyFileVariable: 'SSH_KEY', usernameVariable: 'GIT_USERNAME')]) {
        // Print to check if the variables are set
        sh 'echo SSH key path: $SSH_KEY'
        sh 'echo Git username: $GIT_USERNAME'
        
        // Git clone using the SSH key
        sh '''
            eval `ssh-agent -s`
            ssh-add $SSH_KEY
            git clone git@bitbucket.org:your-repo/project.git
        '''
    }

    
    return false
}
