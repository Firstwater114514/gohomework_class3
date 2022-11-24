# 第三节课作业

## 前排提示：

1、每次关闭postman之类软件页面之前请退出（“/quit”）“账号”。

2、初次使用是没有账号的，需要注册再登录才可使用。

3、使用与留言相关的功能需要注意请求中“comment”的单复数

4、如果使用该服务的用户原本就有以下文件名，就不建议使用

"i"、"login_user"、"n"、"name1"~"name5"、"password1"~"password5"、"comment_1"~"comment6"

## 功能介绍：

### 1、注册

在请求处输入“/register”

key需要username、password

最多可以注册5个账户

### 2、登录

在请求处输入“/login”

key需要username、password

无法多个账户同时登录，换账号前需要退出已登录的账号

登录后才可以使用除了“初始化”、“注册”、“忘记密码”以外的功能

### 3、用令牌获取用户信息

在请求处输入“/user/get”

Header处需要Authorization

### 4、更改密码

在请求处输入“/change password”

key需要new password

更改后会自动退出账号，需要重新登录

### 5、忘记密码

#### 1.获得验证问题

在请求处输入“/forget password”

key需要username

会获得问题，在下一步回答

#### 2.回答问题获得验证码

在请求处输入“/answer”

key需要answer

会获得验证码（checkCode）

#### 3.检查验证码并更改密码

在请求处输入“/check code”

key需要code、username、new password

验证成功密码就更改成功，可以登录了

### 6、留言

在请求处输入“/add comment”

key需要comment

留言板一共可以记录6条留言，满了需要删除才可以发送新的留言

### 7、浏览留言板

在请求处输入“/scan comments”

key不需要任何value

### 8、删除留言

在请求处输入“/delete comment”

key需要num

num是用户想删除的留言的序号

### 9、清空留言板

在请求处输入“/clear comments”

key不需要任何value

### 10、退出账号

在请求处输入“/quit”

key不需要任何value

每次关闭请求之前一定要退出！！！！！！！

### 11、注销账号

在请求处输入“/unsubscribe”

key不需要任何value

### 12、恢复初始化

在请求处输入“/clear all”

key不需要任何value

使用该功能时不可以有账户已登录