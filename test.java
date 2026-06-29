import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;

public class test {

    private static final String API_URL = "http://127.0.0.1:8000/items/";

    public static void main(String[] args) {
        HttpClient client = HttpClient.newBuilder()
                .connectTimeout(Duration.ofSeconds(10))
                .build();

        // JSON payload for the POST request
        String jsonPayload = """
                {
                    "name": "Gaming Laptop",
                    "description": "High performance laptop",
                    "price": 1200.0,
                    "tax": 120.0
                }
                """;

        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_URL))
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(jsonPayload))
                .build();

        try {
            System.out.println("Sending POST request to FastAPI backend in Java...");
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

            System.out.println("HTTP Status Code: " + response.statusCode());
            System.out.println("Response Body:");
            System.out.println(response.body());
        } catch (Exception e) {
            System.err.println("Error making API call: " + e.getMessage());
            e.printStackTrace();
        }
    }
}
