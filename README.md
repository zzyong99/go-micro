# go-micro
Building a simple microservice with golang

## Architecture
- front-end
    - Provided test for connection to broker service
- broker-service
    - Functioned as main entrance of the back-end system
    - Forward auth request to auth-service
- authentication-service
    - Connected to Postgres database to authenticate user login
- logger-service
    - Connected to MongoDB to record log info
- mail-service
    - Use MailHog for email testing 
- listener-service
    - Utilized RabbitMQ to process the request


## Tools
- Back-end Architecture
    - [go-chi](https://github.com/go-chi/chi) simplifies routes building for HTTP services
    - [RabbitMQ](https://github.com/rabbitmq/amqp091-go) buffers and processes the request
    - [gRPC](https://grpc.io/) connects service cross microservice cluster
- Mail Service
    - [MailHog](https://github.com/mailhog/MailHog) tests mail service
    - [go-premailer](https://github.com/vanng822/go-premailer) builds inline styling for HTML mail
    SMTP port usually on 1025 and HTTP port on 8025.