# Run MySQL in Docker
1. docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8.0
2. docker container inspect mysql8 => get IP address "172.17.0.2"
3. docker exec -it mysql8 bash
4. mysql -u root -p => my-secret-pw
5. CREATE DATABASE MYSQLTEST;
6. CREATE USER 'cmis'@'%' IDENTIFIED BY 'Phuongtt@123cmis';
7. GRANT ALL ON MYSQLTEST.* TO 'cmis'@'%';
8. uncomment Mysql section in `cmd/web/main.go`
9. <https://github.com/alochym01/golang-web>
10. (Generate code follow Clean Architechture)<https://sf9v.github.io/posts/generating-the-repository-pattern-in-go/>