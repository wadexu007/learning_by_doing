package main

func TestPizzasHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		input      *Pizzas
		want       string
		statusCode int
	}{
		{
			name:       "without pizzas",
			method:     http.MethodGet,
			input:      &Pizzas{},
			want:       "Error: No pizzas found",
			statusCode: http.StatusNotFound,
		},
		{
			name:   "with pizzas",
			method: http.MethodGet,
			input: &Pizzas{
				Pizza{
					ID:    1,
					Name:  "Foo",
					Price: 10,
				},
			},
			want:       `[{"id":1,"name":"Foo","price":10}]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with bad method",
			method:     http.MethodPost,
			input:      &Pizzas{},
			want:       "Method not allowed",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/orders", nil)
			responseRecorder := httptest.NewRecorder()

			pizzasHandler{tc.input}.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}

func TestOrdersHandler(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		pizzas     *Pizzas
		orders     *Orders
		body       string
		want       string
		statusCode int
	}{
		{
			name:   "with a pizza ID and quantity",
			method: http.MethodPost,
			pizzas: &Pizzas{
				Pizza{
					ID:    1,
					Name:  "Margherita",
					Price: 8,
				},
			},
			orders:     &Orders{},
			body:       `{"pizza_id":1,"quantity":1}`,
			want:       `{"pizza_id":1,"quantity":1,"total":8}`,
			statusCode: http.StatusOK,
		},
		{
			name:   "with a pizza and no quantity",
			method: http.MethodPost,
			pizzas: &Pizzas{
				Pizza{
					ID:    1,
					Name:  "Margherita",
					Price: 8,
				},
			},
			orders:     &Orders{},
			body:       `{"pizza_id":1}`,
			want:       `{"pizza_id":1,"quantity":0,"total":0}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with no pizzas on menu",
			method:     http.MethodPost,
			pizzas:     &Pizzas{},
			orders:     &Orders{},
			body:       `{"pizza_id":1,"quantity":1}`,
			want:       "Error: No pizzas found",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "with GET method and no orders in memory",
			method:     http.MethodGet,
			pizzas:     &Pizzas{},
			orders:     &Orders{},
			body:       "",
			want:       "[]",
			statusCode: http.StatusOK,
		},
		{
			name:   "with GET method and with orders in memory",
			method: http.MethodGet,
			pizzas: &Pizzas{},
			orders: &Orders{
				Order{
					Quantity: 10,
					PizzaID:  1,
					Total:    100,
				},
			},
			body:       "",
			want:       `[{"pizza_id":1,"quantity":10,"total":100}]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "with bad HTTP method",
			method:     http.MethodDelete,
			pizzas:     &Pizzas{},
			orders:     &Orders{},
			body:       "",
			want:       "Method not allowed",
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/orders", strings.NewReader(tc.body))
			responseRecorder := httptest.NewRecorder()

			handler := ordersHandler{tc.pizzas, tc.orders}
			handler.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
