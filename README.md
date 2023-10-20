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
    - [go-chi](github.com/go-chi/chi), to simplify routes building for HTTp services
- Mail Service
    - [MailHog](https://github.com/mailhog/MailHog), to test mail service
    - [go-premailer](github.com/vanng822/go-premailer), to build inline styling for HTML mail
SMTP port on 1025 and HTTP port on 8025