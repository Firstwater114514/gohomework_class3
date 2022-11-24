package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/api/middleware"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/dao"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/model"
	"lanshan_homework/go1.19.2/go_homework/class_3_work_lv1/utils"
	"strconv"
	"time"
)

var (
	n     int
	i     int
	kindC int8
	kind  int8
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "部分数据未输入，请检查")
		return
	}
	nn, _ := ioutil.ReadFile("n") ////////////////////////////////提取n
	var errN error
	n, errN = strconv.Atoi(string(nn))
	if errN != nil {
		panic(errN)
	} ///////////////////////////////////////////////提取n
	if n >= 6 {
		utils.RespFail(c, "注册人数已达上限~非常抱歉")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		utils.RespFail(c, "用户名已被使用")
		return
	}
	dao.AddUser(username, password)
	n++
	nnu := strconv.Itoa(n)
	nb := []byte(nnu)
	errs := ioutil.WriteFile("n", nb, 0644)
	if errs != nil {
		panic(errs)
	}
	utils.RespSuccess(c, "注册成功!快去登录吧~")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	checkUser, _ := ioutil.ReadFile("login_user")
	if string(checkUser) != "" {
		utils.RespFail(c, "请先退出当前帐号！")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "该用户不存在")
		return
	}
	selectPassword := dao.SelectPasswordFromUsername(username)
	if selectPassword != password {
		utils.RespFail(c, "密码错误！")
		return
	}
	user := []byte(username)
	errU := ioutil.WriteFile("login_user", user, 0644)
	if errU != nil {
		panic(errU)
	}
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "sqy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.LoginSuccess(c, tokenString, "将请求中的login改成add comment并且输入comment就可以写留言啦", "将请求中的login改成delete comment并且输入想删除的留言序号num", "将请求中的login改成change password并且输入new password可以更改密码哦", "注销账号请使用unsubscribe", "记得关闭窗口前一定要退出账号哦!!!!!用quit")
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func changePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Change{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	password := dao.SelectPasswordFromUsername(username)
	newPassword := c.PostForm("new password")
	if password == newPassword {
		utils.RespFail(c, "新密码与旧密码相同~")
		return
	}
	dao.ChangePassword(username, newPassword)
	dao.Quit()
	utils.RespSuccess(c, "更改密码成功，请重新登录~")
	return
}
func forgetPassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Forget{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "该用户不存在")
		return
	}
	selectPassword := dao.SelectPasswordFromUsername(username)
	if selectPassword == "" {
		utils.RespFail(c, "该用户不存在")
		return
	}
	utils.Question(c, "问题：你的大学叫什么？(提示: CQUPT)", "把请求中的forget password改成answer并且输入answer")
}
func answer(c *gin.Context) {
	if err := c.ShouldBind(&model.Question{}); err != nil {
		utils.RespFail(c, "请输入答案")
		return
	}
	answer := c.PostForm("answer")
	check := dao.CheckAnswer(answer)
	if !check {
		utils.RespFail(c, "验证码不正确，请重试")
		return
	}
	maxNum := 999999
	kindC = 2
	utils.AnswerRight(c, "验证成功！", "将请求中的answer改成check code并且输入username,code,new password", "checkCode是验证码", dao.MakeCode(maxNum))
}
func checkCode(c *gin.Context) {
	if err := c.ShouldBind(&model.Code{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	code := c.PostForm("code")
	kind = 2
	if kindC != 2 {
		utils.RespFail(c, "请先获取验证码")
		return
	}
	check := dao.CheckCode(code, kindC, kind)
	if !check {
		utils.RespFail(c, "验证码错误！请重新输入")
		return
	}
	username := c.PostForm("username")
	newPassword := c.PostForm("new password")
	dao.ChangePassword(username, newPassword)
	utils.RespSuccess(c, "更改密码成功！请重新登录哦")
	return
}
func addComment(c *gin.Context) {
	if err := c.ShouldBind(&model.AddComment{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	comment := c.PostForm("comment")
	in, _ := ioutil.ReadFile("i") ////////////////////////////////提取i
	var errI error
	i, errI = strconv.Atoi(string(in))
	if errI != nil {
		panic(errI)
	} ///////////////////////////////////////////////提取i
	if i >= 7 {
		utils.RespFail(c, "留言已达数量上限")
		return
	}
	dao.Refresh(i)
	dao.AddComment(i, username, comment)
	i++
	inu := strconv.Itoa(i)
	ib := []byte(inu)
	errs := ioutil.WriteFile("i", ib, 0644)
	if errs != nil {
		panic(errs)
	}
	utils.RespSuccess(c, "留言成功，把请求中的add comment改成scan comments查看留言板")
}
func scanComments(c *gin.Context) {
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	in, _ := ioutil.ReadFile("i")
	var errI error
	i, errI = strconv.Atoi(string(in))
	if errI != nil {
		panic(errI)
	}
	if i == 1 {
		utils.RespSuccess(c, "还没有人留言哦，你来发表第一个留言吧~")
		return
	}
	dao.Refresh(i)
	var read [6]string
	for j := 1; j <= 6; j++ {
		jn := strconv.Itoa(j)
		filename := "comment_" + jn
		scan, _ := ioutil.ReadFile(filename)
		read[j-1] = string(scan)
	}
	utils.Comments(c, "留言板：", read[0], read[1], read[2], read[3], read[4], read[5])
}
func deleteComment(c *gin.Context) {
	if err := c.ShouldBind(&model.DeleteComment{}); err != nil {
		utils.RespFail(c, "请输入想删除的留言序号哦")
		return
	}
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	num := c.PostForm("num")
	in, _ := ioutil.ReadFile("i")
	var errI error
	i, errI = strconv.Atoi(string(in))
	if errI != nil {
		panic(errI)
	}
	if i == 1 {
		utils.RespSuccess(c, "还没有人留言哦，没有留言可以删呢~")
		return
	}
	numI, errN := strconv.Atoi(num)
	if errN != nil {
		panic(errN)
	}
	if numI > i-1 {
		utils.RespFail(c, "没有该序号的留言")
		return
	}
	dao.Refresh(i)
	fault := 1
	dao.DeleteComment(numI, fault)
	if fault == 0 {
		utils.RespFail(c, "没有该序号的留言")
		return
	}
	i--
	inu := strconv.Itoa(i)
	ib := []byte(inu)
	errs := ioutil.WriteFile("i", ib, 0644)
	if errs != nil {
		panic(errs)
	}
	if i == 1 {
		utils.RespSuccess(c, "成功删除留言，留言板暂无留言")
		return
	}
	var read [6]string
	for j := 1; j <= 6; j++ {
		jn := strconv.Itoa(j)
		filename := "comment_" + jn
		scan, _ := ioutil.ReadFile(filename)
		read[j-1] = string(scan)
	}
	utils.Comments(c, "删除留言成功，留言板：", read[0], read[1], read[2], read[3], read[4], read[5])
}
func clearComments(c *gin.Context) {
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	in, _ := ioutil.ReadFile("i")
	var errI error
	i, errI = strconv.Atoi(string(in))
	if errI != nil {
		panic(errI)
	}
	if i == 1 {
		utils.RespFail(c, "留言板本来就是空的哦~")
		return
	}
	dao.ClearComments()
	i = 1
	inu := strconv.Itoa(i)
	ib := []byte(inu)
	errs := ioutil.WriteFile("i", ib, 0644)
	if errs != nil {
		panic(errs)
	}
	utils.RespSuccess(c, "清除留言板成功")
}
func quit(c *gin.Context) {
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else {
		dao.Quit()
	}
	utils.RespSuccess(c, "成功退出账号")
	return
}
func unsubscribe(c *gin.Context) {
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username == "" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	}
	dao.Unsubscribe(username)
	n--
	nnu := strconv.Itoa(n)
	nb := []byte(nnu)
	errs := ioutil.WriteFile("n", nb, 0644)
	if errs != nil {
		panic(errs)
	}
	dao.Quit()
	utils.RespSuccess(c, "注销账户成功")
}
func clearAll(c *gin.Context) {
	user, _ := ioutil.ReadFile("login_user")
	username := string(user)
	if username != "" {
		utils.RespFail(c, "该功能仅退出账号后可使用")
		return
	}
	dao.ClearComments()
	i = 1
	inu := strconv.Itoa(i)
	ib := []byte(inu)
	errI := ioutil.WriteFile("i", ib, 0644)
	if errI != nil {
		panic(errI)
	}
	dao.ClearData()
	n = 1
	nnu := strconv.Itoa(n)
	nb := []byte(nnu)
	errN := ioutil.WriteFile("n", nb, 0644)
	if errN != nil {
		panic(errN)
	}
	utils.RespSuccess(c, "成功初始化该系统！")
}
