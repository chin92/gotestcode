package your.package;

import org.apache.http.HttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.junit.Assert;
import org.junit.Test;

import java.io.IOException;

public class YourIntegrationTest {

    @Test
    public void testEndpointReturns200() throws IOException {
        // Replace with your actual endpoint URL
        String url = "https://your-endpoint.com/api/resource";

        // Create an HTTP client and make the request
        try (CloseableHttpClient httpClient = HttpClients.createDefault()) {
            HttpGet request = new HttpGet(url);
            HttpResponse response = httpClient.execute(request);

            // Check if the status code is 200
            int statusCode = response.getStatusLine().getStatusCode();
            Assert.assertEquals(200, statusCode);
        }
    }
}
