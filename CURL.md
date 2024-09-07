# health
curl -i -X GET http://localhost:8091/health

# create user
curl -i -X PUT -H 'Content-Type: application/json' -d '{"first_name": "FirstName", "last_name": "LastName", "pswd_hash": "$2a$10$wrfLakZMCZHQStxyvmfmWuIF8ovj2Tcbdc9tH3VEf8MPWntBLg55W", "email": "a@a.com", "country": "UK"}' http://localhost:8091/users

# get user
curl -i -X GET http://localhost:8091/users/3746e0e3-d4d8-4ad1-8099-f5a3b5ab9a6d

# list user
curl -i -X GET 'http://localhost:8091/users?page-number=1&page-size=10&country=UK,USA'

# update user
curl -i -X POST -H 'Content-Type: application/json' -d '{"first_name": "NewFirstName", "email": "new@a.com"}' http://localhost:8091/users/3746e0e3-d4d8-4ad1-8099-f5a3b5ab9a6d

# delete user
curl -i -X DELETE http://localhost:8091/users/e26ecf40-3a3d-4223-aac8-2f955440a618
