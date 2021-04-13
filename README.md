# Auth service
An auth service that issue the JWT-token based on the auth_code.  
A user requests a code by `POST /code` request and has to pass it to 
the `POST /token` endpoint.  
# Packages
* [API](https://github.com/cegorah/auth_service/tree/master/api)  
The package that is responsible for serving the API  
* [AuthCode](https://github.com/cegorah/auth_service/tree/master/pkg/auth_code)  
Generate and check a random `int` code based on provided length.  
The code will be stored by [Cacher](https://github.com/cegorah/auth_service/blob/master/pkg/auth_code/code_validator.go#L16) 
for further checking.  
* [MsgSend](https://github.com/cegorah/auth_service/tree/master/pkg/msg_send)  
Using for sending code by email or SMS.
* [Authorization](https://github.com/cegorah/auth_service/tree/master/pkg/authorization)  
Using for creating users and issue JWE-tokens. 
* [Cache](https://github.com/cegorah/auth_service/tree/master/internal/cache)  
Is implementation of [Cacher](https://github.com/cegorah/auth_service/blob/master/pkg/auth_code/code_validator.go#L16)
interface that is using by [CodeValidator](https://github.com/cegorah/auth_service/blob/master/pkg/auth_code/code_validator.go#L27) 
for storing and getting a validation code.
