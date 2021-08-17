# gRPC Demo

    1. Get the code:
    ```console
    $ go get github.com/chyidl/begin-go-micro/examples/go-grpc/demo/server/user_server
    $ go get github.com/chyidl/begin-go-micro/examples/go-grpc/demo/client/user_client
    ```
    
    2. Run the server:
    ```console
    $ $(go env GOPATH)/bin/user_server &
    ```
    
    3. Run the client:
    ```console
    $ $(go env GOPATH)/bin/user_client
    ```


## Create a Self-Signed SSL Certificate
```markdown
Step 1 - Creating the SSL Certificate
    TLS/SSL works by using a combination of a public certificate and a private key.
        The SSL Key: is kept secret on the server [用于加密传输给客户端的内容]
        The SSL Certificate: is publicly shared [证书公开，任何人可以获取，用于解密]




$ sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./certs/selfsigned.key -out ./certs/selfsigned.crt
# 生成server.key 私钥文件 -- 密码1234
$ sudo openssl genrsa -des3 -out ./certs/server.key 2048  

# 生成server.csr 证书请求
$ sudo openssl req -new -key ./certs/server.key -out ./certs/server.csr

# 去除密码
$ sudo openssl rsa -in server.key -out server_no_passwd.key

# 生成服务端证书
$ sudo openssl x509 -req -days 265 -in ./certs/server.csr -signkey ./certs/server_no_passwd.key -out ./certs/server.crt
```