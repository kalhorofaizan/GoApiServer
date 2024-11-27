## Steps for setup

- mkdir keys
- openssl genpkey -algorithm ed25519 -outform PEM -out keys/private.pem && openssl pkey -in keys/private.pem -pubout -out keys/public.pem // it will create public and private key for jwt


## Feature
- basic Route
- Body validator
- Auth validator
- Jwt Token



## Library used
- github.com/golang-jwt/jwt/v5
- github.com/go-playground/validator/v10
- github.com/joho/godotenv
- net/http