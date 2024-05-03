# heilbronn-mensa-api
Webscraping API for Bildungscampus Heilbronn Mensa by jkauker

### Sample Response for `/api/menu/today`
```json
"menu": {
  "dishes": [
    {
      "name": "Sample Dish",
      "price": {
        "student": "3,20"
        "guest": "6,76",
        "worker": "4,99"
      },
      type: "Dessert"
    },
    {
      "name": "Sample Dish 2",
      "price": {
        "student": "1,30"
        "guest": "5,49",
        "worker": "3,45"
      },
      type: "Soup"
    }
  ],
  "lang": "de"
},
date: "03.05.2024"
```

### Dish Types:
- dessert
- dessert-vegan
- soup
- meat
- vegan
- vegetarian

