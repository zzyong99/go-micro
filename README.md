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
- logger-service
    - Connected to MongoDB to record log info
- mail-service
    - Use MailHog for email testing 


## Tools
- Back-end Service
    - [go-chi](https://github.com/go-chi/chi) simplifies routes building for HTTP services
- Mail Service
    - [MailHog](https://github.com/mailhog/MailHog) tests mail service
    - [go-premailer](https://github.com/vanng822/go-premailer) builds inline styling for HTML mail
    SMTP port usually on 1025 and HTTP port on 8025.