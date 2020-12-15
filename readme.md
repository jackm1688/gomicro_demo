运行方法:<br>
cd service<br>
go run srv_main.go --registry consul --registry_address 192.168.0.50:8500<br>

cd client<br>
go run cliient_main.go  --registry consul --registry_address 192.168.0.52:8500<br>
ret: 你好,柠檬酱,欢迎到 西安<br>
ret: 301<br>

cd api<br>
go run api.go  --registry consul --registry_address 192.168.0.51:8500<br>

microv113 -version  <br>
micro version v1.13.1-6d09ae2-1571469584<br>

# 启动命令--API网关<br>
microv113 --registry consul --registry_address 192.168.0.52:8500 api --handler=api<br>

api测试<br>
curl http://127.0.0.1:8080/Hello/Hello/add?num1=11&num2=100<br>
curl http://127.0.0.1:8080/Hello/Hello/SayHi?name=jackm&adrress=西安<br>


