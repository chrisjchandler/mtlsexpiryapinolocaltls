Curl Example
curl --cert client.crt --key client.key --cacert ca.crt -X PUT https://localhost:8443/upload

API is a push API, reads in Certs so scripting to routinely check is run client side

Output example

HTTP/1.1 200 OK
Content-Type: application/json

{"common_name":"example.com","days_until_expiry":30}

