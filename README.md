A mini project for social network service's user management module, consisting of user registration, user login, logout and user's follow of another user. 

This project is a simple practice for the use of gin-gonic webframe and grpc service.To run this project on your local machine, follow these steps: 

1. Replace the relative configurations in `minisns/config/config.go` by your own settings. Ensure you have create the database and tables in your MySQL server. MySQL tables structure are contained in the `.sql` files under the `deploy` directory.

2. Make sure the `TCP_PORT` and `HTTP_PORT` on your machine is not in use, and then execute the following commands:
    ```
    cd your_path/minisns
    sudo go run cmd/tcpserver/main.go
    sudo go run cmd/httpserver/main.go
    ```
3. After step2, you can see the tcpserver and httpserver are both running in your local machine, and you can try to access all the api now! Refer to the `api_access_demo` file for convenience.
