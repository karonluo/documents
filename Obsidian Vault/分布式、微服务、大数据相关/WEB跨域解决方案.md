#编程语言/Java语言 #编程语言/GO语言 #编程语言/Python语言 #JavaScript #DEV #OPS #WEB/WEB后端/框架 
# WEB 跨域解决方案
## 描述

目前 Web 应用 B/S 架构，常采用 前后端分离方案，势必会遇到大量跨域访问的问题。
按照本文所述配置则可以完美解决问题。
#编程语言/JavaScript语言 
## 前端 JQuery 或 Axios 发送请求配置请求头

基于 axios
~~~JavaScript
function login() {
	var data = JSON.stringify({
		"login_name": "admin",
		"password": "pass@@word123",
		"code": "53"
		});
		
		var config = {
		method: 'post',
		url: 'http://192.168.1.100:8080/login',
		headers: { 
			"Context-Type":"application/x-www-form-urlencoded"
		},
		data : data
		};
		axios(config)
		.then(function (response) {
		console.log(JSON.stringify(response.data));
		})
		.catch(function (error) {
		console.log(error);
		});
}
// Header 设置 Context-Type 是关键，如果是 application/json 方式，需要2次请求后端,造成稳定性和性能问题，第一次是OPTIONS 第二次才是其他类型的
~~~
基于 jquery
~~~JavaScript
function login() {
	var data = JSON.stringify({
		"login_name": "admin",
		"password": "pass@@word123",
		"code": "53"
	});
	var settings = {
		"url": "http://192.168.1.100:8080/login",
		"method": "POST",
		"headers": {
			"Context-Type":"application/x-www-form-urlencoded"
		},
		"data": data
	};

	$.ajax(settings).done(function (response) {
		console.log(response);
	});
}
~~~
## 后端 Web 服务
#WEB/WEB后端/框架/GO语言/IRIS
### Go Iris 框架 相关代码

拦截器代码
~~~Go
// 拦截器
func UseBefore(ctx iris.Context){
	// 处理第一次用于跨域的验证请求
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.StatusCode(204)
		return
	}
	// 处理后续的真实的请求
	// 代码略过
}
~~~

服务或页面注册绑定代码
~~~Go
func Cros(ctx iris.Context){
	ctx.Text("")
}
func main(){
	app:=iris.New()
	app.Use(UseBefore)
	app.Get("/login", WSLogin)
	app.Options("/login", Cros)
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
// 此时前端页面 jquery 或 axios 调用 /login 能够顺利调用了。
~~~

#WEB/WEB后端/框架/Java语言/SpringBoot
### Java SpringBoot 相关代码

#WEB/WEB后端/框架/Python语言/Django
### Python Django 相关代码

#WEB/容器/Nginx 
### Nginx 配置


### 建议

前端页面：Header 设置 Context-Type 是关键，如果是 application/json 方式，需要2次请求后端,造成稳定性和性能问题，第一次是 OPTIONS 第二次才是其他类型的请求，因此最好设置成 application/x-www-form-urlencoded 减少访问次数，同时后端也可以减少 Options 的方法绑定。