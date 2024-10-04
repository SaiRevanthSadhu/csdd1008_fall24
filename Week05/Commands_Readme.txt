1. To create the item :- curl -X POST -d "{\"model\":\"Defender\",\"year\":2021}" -H "Content-Type: application/json" http://localhost:8080/cars
-- created the second car :- curl -X POST -d "{\"model\":\"Tata Nexon\",\"year\":2022}" -H "Content-Type: application/json" http://localhost:8080/cars
2. To list the items :- curl http://localhost:8080/cars
3. To update the items :- curl -X PUT -d "{\"model\":\"Range Rover\",\"year\":2022}" -H "Content-Type: application/json" http://localhost:8080/cars/1
4. To delete the items :- curl -X DELETE http://localhost:8080/cars/1

Attached the screenshot for the execution of the cars created, updated, read and deleted. 