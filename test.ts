// TypeScript interfaces representing the FastAPI backend models

export interface Item {
  name: string;
  description?: string;
  price: number;
  tax?: number;
}

export interface ItemResponse {
  message: string;
  item: Item;
  total_price: number;
}

const API_BASE_URL = "http://127.0.0.1:8000";

/**
 * Function to make a POST request to the FastAPI /items/ endpoint
 */
export async function createItem(itemData: Item): Promise<ItemResponse> {
  try {
    const response = await fetch(`${API_BASE_URL}/items/`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(itemData),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(`HTTP Error ${response.status}: ${JSON.stringify(errorData)}`);
    }

    const result: ItemResponse = (await response.json()) as ItemResponse;
    return result;
  } catch (error) {
    console.error("Failed to create item:", error);
    throw error;
  }
}

// Example execution
async function main() {
  const newItem: Item = {
    name: "Gaming Laptop",
    description: "High performance laptop",
    price: 1200.0,
    tax: 120.0,
  };

  console.log("Sending POST request to FastAPI backend...");
  const response = await createItem(newItem);
  console.log("API Response:", response);
}

// Execute main function if run directly in Node/Deno/Bun environment
if (typeof process !== "undefined" && require.main === module) {
  main();
}
