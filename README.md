### Task Experimental

Hello! I'm glad to review your task. The project has been developed using the Go programming language, and efforts have
been made to ensure that the project is modular. For testing, some data has been created in the database.

### how deploy:

````
docker compose up -d
````

URL:[localhost:8787]()

If you need to make changes, you can refer to the ``config.yml`` for information such as database connection details.

### API

1- To announce a delay in an order, this request is used.

````
curl  -X POST --location 'http://localhost:8787/delay_report' \
--header 'Content-Type: application/json' \
--data '{
    "order_id": 1
}'
````

2- The orders in the delay queue are assigned to the employee through the API on a first-come, first-served basis.

```
curl --location 'http://localhost:8787/delay_report/1'
```

3- Write a query for this API that returns the delay values for each store in the past week, sorted in descending order.

```
curl --location 'http://localhost:8787/delay_report/last_week'
```