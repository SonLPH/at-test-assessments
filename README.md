### AT HOME ASSESSMENT ### 
1. Clone source code
```shell
git clone https://github.com/SonLPH/at-test-assessments.git
```
2. Create .env file if not exists in root directory:
```shell
cp .env.example .env
```
3. Run docker to build database up
```shell
make docker-compose-dev-up
```
4. Open 3 terminal for 3 services
```go
go run services/bookingservice/main.go
```
```go
go run services/pricingservice/main.go
```
```go
go run services/sendservice/main.go
```

5. Booking service host in localhost:4444, Pricing service host in localhost:4445 and Send service host in localhost:4446. So first thing first make a call to mock employee in database.
```shell
curl localhost:4446/mock
```

6. Now use POSTMAN for using API. API in this repo I have provided:
+  POST API Create Booking - localhost:4444/booking
    
    request example
    
    {
        "of_user_id" : "3",
        "booking_total" : "100",
        "booking_type" : "house"
    }

+  GET API Pricing - localhost:4445/pricing

    request example
    {
        "day" : "2024-04-07",
        "employee_id" : "0"
    }