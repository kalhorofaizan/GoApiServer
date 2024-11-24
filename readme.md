## Steps for setup

- mkdir keys
- openssl genpkey -algorithm ed25519 -outform PEM -out keys/private.pem && openssl pkey -in keys/private.pem -pubout -out keys/public.pem // it will create public and private key for jwt