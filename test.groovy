def version1 = '0.7'
def version2 = '0.0.6'

// Function to compare version strings
boolean isVersionLessThan(String v1, String v2) {
    def v1Parts = v1.tokenize('.')
    def v2Parts = v2.tokenize('.')

    def maxLength = Math.max(v1Parts.size(), v2Parts.size())

    for (int i = 0; i < maxLength; i++) {
        def v1Part = (i < v1Parts.size()) ? v1Parts[i].toInteger() : 0
        def v2Part = (i < v2Parts.size()) ? v2Parts[i].toInteger() : 0

        if (v1Part < v2Part) {
            return true
        } else if (v1Part > v2Part) {
            return false
        }
    }
    return false
}

if (isVersionLessThan(version2, version1)) {
    def command = 'curl http://example.com'
    def process = command.execute()
    process.waitFor()
    println process.text
} else {
    println "${version2} is not less than ${version1}"
}
