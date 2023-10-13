# go-micro
Building a simple microservoce with golang

## Architecture
- front-end
    - Provided test for connection to broker service
- broker-service
    - Functioned as main entrance of the back-end system
    - Forward auth request to auth-service
- authentication-service
    - Connected to Postgres database to authenticate user login