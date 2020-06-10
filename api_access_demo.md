1. register

curl -H 'Content-Type: application/json' -d '{"username":"jefferey", "nickname":"jefferey", "password":"jefferey","email":"jefferey@email.com"}' http://localhost:10001/register

----------------------------------------------------------------------

2. login

curl -H 'Content-Type: application/json' -d '{"username":"jefferey", "password":"jefferey"}' http://localhost:10001/login

----------------------------------------------------------------------

3. follow

curl -b "identity=e0c0fe67-d9d6-4a72-bfe5-fa7b84199ac8" -H 'Content-Type: application/json' -d '{"followee":388772787}' http://localhost:10001/api/follow

----------------------------------------------------------------------

4. followlist

curl -b "identity=e0c0fe67-d9d6-4a72-bfe5-fa7b84199ac8" "http://localhost:10001/api/follow_list?pageno=1&count=20"

----------------------------------------------------------------------

5. edit_user

curl -b "identity=4310c702-5db7-44fb-8884-1cc1d5ebf0c9" -H 'Content-Type: application/json' -d '{"nickname":"jefferey","email":"jefferey@email.com","password":"jefferey"}' http://localhost:10001/api/edit_user

----------------------------------------------------------------------

6. logout

curl -b "identity=e0c0fe67-d9d6-4a72-bfe5-fa7b84199ac8" -XPOST "http://localhost:10001/api/logout"

----------------------------------------------------------------------

stress test
go-wrk -c=1 -t=1 -n=1 "http://localhost:10001/api/follow_list?pageno=1&count=20"