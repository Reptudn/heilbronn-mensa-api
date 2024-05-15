# Heilbronn Mensa API
Webscraping API for Bildungscampus Heilbronn Mensa by [jkauker](https://profile.intra.42.fr/users/jkauker)

### How to run?
1. Need to have docker installed
2. Be able to run Makefiles (usually preinstalled on linux systems)
3. run `make run`

### Route
#### /api/mensa/menu/lang/date
returns the menu for that day

### Parameters
- lang: `de`/`en`
- date format: `dd-mm-yyyy`

### Sample Response for `/api/mensa/menu/de/today`
```json
"menu": {
  "dishes": [
    {
      "name": "Sample Dish",
      "price": {
        "student": "3,20",
        "guest": "6,76",
        "worker": "4,99"
      },
      "type": "Dessert"
    },
    {
      "name": "Sample Dish 2",
      "price": {
        "student": "1,30",
        "guest": "5,49",
        "worker": "3,45"
      },
      "type": "Soup"
    }
  ],
  "lang": "de"
},
"date": "03.05.2024"
```

### Dish Types:
- dessert
- dessert-vegan
- soup
- meat
- vegan
- vegetarian
