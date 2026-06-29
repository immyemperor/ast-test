import json
import urllib.request
import urllib.error

API_URL = "http://127.0.0.1:8000/items/"

def create_item():
    payload = {
        "name": "Gaming Laptop",
        "description": "High performance laptop",
        "price": 1200.0,
        "tax": 120.0
    }

    headers = {
        "Content-Type": "application/json"
    }

    data = json.dumps(payload).encode("utf-8")
    req = urllib.request.Request(API_URL, data=data, headers=headers, method="POST")

    print("Sending POST request to FastAPI backend in Python...")
    try:
        with urllib.request.urlopen(req) as response:
            status_code = response.getcode()
            response_bytes = response.read()
            response_json = json.loads(response_bytes.decode("utf-8"))

            print(f"HTTP Status Code: {status_code}")
            print("API Response:")
            print(json.dumps(response_json, indent=2))
            return response_json
    except urllib.error.HTTPError as e:
        error_content = e.read().decode("utf-8")
        print(f"HTTP Error {e.code}: {error_content}")
    except urllib.error.URLError as e:
        print(f"URL Error (Is server running?): {e.reason}")

if __name__ == "__main__":
    create_item()
