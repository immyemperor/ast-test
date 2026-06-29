from fastapi import FastAPI
from pydantic import BaseModel, Field

app = FastAPI()

# Enhanced Pydantic model for Item with validations and descriptions
class Item(BaseModel):
    name: str = Field(..., min_length=1, max_length=100, description="Name of the item", json_schema_extra={"example": "Gaming Laptop"})
    description: str | None = Field(default=None, max_length=300, description="Optional description")
    price: float = Field(..., gt=0, description="Price must be greater than zero", json_schema_extra={"example": 999.99})
    tax: float | None = Field(default=None, ge=0, description="Tax amount", json_schema_extra={"example": 50.00})

# Response model schema
class ItemResponse(BaseModel):
    message: str
    item: Item
    total_price: float
    qty: int
    total_qty: int
    date: str

@app.get("/")
def read_root():
    return {"message": "Welcome to the FastAPI App!"}

# POST request endpoint returning structured ItemResponse
@app.post("/items/", response_model=ItemResponse)
def create_item(item: Item):
    total_price = item.price + (item.tax if item.tax else 0.0)
    
    return ItemResponse(
        message="Item created successfully!",
        item=item,
        total_price=round(total_price, 2)
    )

