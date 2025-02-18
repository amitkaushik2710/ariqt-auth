# ariqt-auth

You can also run the swagger document by
http://localhost:7001/swagger/index.html#/

1. Run Microservice:

go run main.go

2. Sign up API:

curl -X 'POST' \
  'http://localhost:7001/api/v1/signup' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "a",
  "password": "a"
}'

3. Sign in API:

curl -X 'POST' \
  'http://localhost:7001/api/v1/signin' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "a",
  "password": "a"
}'

4. Test Req Data (Authorization of token) API:

curl -X 'POST' \
  'http://localhost:7001/api/v1/test-req' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImIiLCJleHAiOjE3Mzk5MTUwNTF9.x_emwIAFEeCvCtmx42d3xmqM0Xsgi-WV9tTwyVeXASg' \
  -H 'Content-Type: application/json' \
  -d '{
  "data": true
}'

> [!CAUTION]
> update your auth token in Authorization after the Bearer

5. Refresh Token (Mechanism to refresh a token) API:

curl -X 'GET' \
  'http://localhost:7001/api/v1/refresh-token' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImEiLCJleHAiOjE3Mzk5MTUzNDN9.H972O0DpIWWIyQhexSspqs_kTv4fJmDvlsz51E1zaQs'

> [!CAUTION]
> update your auth token in Authorization after the Bearer

6. Revoke Token (Revocation of token) API:

curl -X 'DELETE' \
  'http://localhost:7001/api/v1/revoke-token' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImIiLCJleHAiOjE3Mzk5MTUwNTF9.x_emwIAFEeCvCtmx42d3xmqM0Xsgi-WV9tTwyVeXASg'

> [!CAUTION]
> update your auth token in Authorization after the Bearer