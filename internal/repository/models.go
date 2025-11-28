package repository

type Order struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"cust_id"`
	CreatedAt  string `json:"created_at"`
}

type OrderItem struct {
	ID        int64   `json:"id"`
	OrderID   int64   `json:"order_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float32 `json:"price"`
}

type MenuItem struct {
    ID              int64    `json:"id"`
    Name            string   `json:"name"`
    Description     string   `json:"description"`
    Category        string   `json:"category"`
    BasePrice       float32  `json:"base_price"`
    ImageURL        string   `json:"image_url"`
    IsAvailable     bool     `json:"is_available"`
    Customizations  []int64  `json:"customizations"`
}

type Menu struct {
    MenuItems       []MenuItem `json:"menu_items"`
    Customizations  []interface{} `json:"customizations"`
    Categories      []interface{} `json:"categories"`
}