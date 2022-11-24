package dao

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

var (
	name1, _     = ioutil.ReadFile("name1")
	password1, _ = ioutil.ReadFile("password1")
	name2, _     = ioutil.ReadFile("name2")
	password2, _ = ioutil.ReadFile("password2")
	name3, _     = ioutil.ReadFile("name3")
	password3, _ = ioutil.ReadFile("password3")
	name4, _     = ioutil.ReadFile("name4")
	password4, _ = ioutil.ReadFile("password4")
	name5, _     = ioutil.ReadFile("name5")
	password5, _ = ioutil.ReadFile("password5")
	com          [6]string
)
var database = map[string]string{
	string(name1): string(password1),
	string(name2): string(password2),
	string(name3): string(password3),
	string(name4): string(password4),
	string(name5): string(password5),
}

func AddUser(username, password string) {
	for j := 1; j <= 6; j++ {
		jn := strconv.Itoa(j)
		user := "name" + jn
		readN, _ := ioutil.ReadFile(user)
		if string(readN) == "" {
			data1 := []byte(username)
			data2 := []byte(password)
			adduser := "name" + jn
			addPassword := "password" + jn
			err1 := ioutil.WriteFile(adduser, data1, 0644)
			if err1 != nil {
				panic(err1)
			}
			err2 := ioutil.WriteFile(addPassword, data2, 0644)
			if err2 != nil {
				panic(err2)
			}
			database[username] = password
			break
		}
	}
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}
func ChangePassword(username, newPassword string) {
	database[username] = newPassword
	for j := 1; j <= 6; j++ {
		js := strconv.Itoa(j)
		filename := "name" + js
		find, _ := ioutil.ReadFile(filename)
		if username == string(find) {
			passwordFile := "password" + js
			change := []byte(newPassword)
			err := ioutil.WriteFile(passwordFile, change, 0644)
			if err != nil {
				panic(err)
			}
			break
		}
	}
}
func CheckAnswer(answer string) bool {
	if answer == "CQUPT" {
		return true
	}
	return false
}
func MakeCode(maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	sn := strconv.Itoa(secretNumber)
	dataCode := []byte(sn)
	err1 := ioutil.WriteFile("checkCode", dataCode, 0644)
	if err1 != nil {
		panic(err1)
	}
	return secretNumber
}
func CheckCode(code string, kindC, kind int8) bool {
	checkCode, _ := ioutil.ReadFile("checkCode")
	if code == string(checkCode) && kindC == kind {
		dataCode := []byte("-1")
		err1 := ioutil.WriteFile("checkCode", dataCode, 0644)
		if err1 != nil {
			panic(err1)
		}
		return true
	}
	return false
}
func AddComment(i int, username, comment string) {
	in := strconv.Itoa(i)
	com[i-1] = username + ":" + comment
	comm := []byte(com[i-1])
	filename := "comment_" + in
	errC := ioutil.WriteFile(filename, comm, 0644)
	if errC != nil {
		panic(errC)
	}
}
func Refresh(i int) {
	if i > 1 {
		for j := 1; j <= i-1; j++ {
			jn := strconv.Itoa(j)
			filename := "comment_" + jn
			read, _ := ioutil.ReadFile(filename)
			com[j-1] = string(read)
		}
	}
}
func Start() {
	in, _ := ioutil.ReadFile("i")
	var errI error
	_, errI = strconv.Atoi(string(in))
	if errI != nil {
		_1 := "1"
		_i := []byte(_1)
		err := ioutil.WriteFile("i", _i, 0644)
		if err != nil {
			panic(err)
		}
	}
	nn, _ := ioutil.ReadFile("n")
	var errN error
	_, errN = strconv.Atoi(string(nn))
	if errN != nil {
		_1 := "1"
		_n := []byte(_1)
		err := ioutil.WriteFile("n", _n, 0644)
		if err != nil {
			panic(err)
		}
	}
}
func DeleteComment(numI, fault int) {
	switch numI {
	case 1:
		com[0] = com[1]
		com[1] = com[2]
		com[2] = com[3]
		com[3] = com[4]
		com[4] = com[5]
		com[5] = ""
	case 2:
		com[1] = com[2]
		com[2] = com[3]
		com[3] = com[4]
		com[4] = com[5]
		com[5] = ""
	case 3:
		com[2] = com[3]
		com[3] = com[4]
		com[4] = com[5]
		com[5] = ""
	case 4:
		com[3] = com[4]
		com[4] = com[5]
		com[5] = ""
	case 5:
		com[4] = com[5]
		com[5] = ""
	case 6:
		com[5] = ""
	default:
		fault = 0
	}
	for j := numI; j <= 6; j++ {
		in := strconv.Itoa(j)
		comm := []byte(com[j-1])
		filename := "comment_" + in
		errC := ioutil.WriteFile(filename, comm, 0644)
		if errC != nil {
			panic(errC)
		}
	}
}
func ClearComments() {
	for j := 1; j <= 6; j++ {
		in := strconv.Itoa(j)
		comm := []byte("")
		filename := "comment_" + in
		errC := ioutil.WriteFile(filename, comm, 0644)
		if errC != nil {
			panic(errC)
		}
	}
	for j := 1; j <= 6; j++ {
		jn := strconv.Itoa(j)
		filename := "comment_" + jn
		read, _ := ioutil.ReadFile(filename)
		com[j-1] = string(read)
	}
}
func Quit() {
	quit := []byte("")
	clear := ioutil.WriteFile("login_user", quit, 0644)
	if clear != nil {
		panic(clear)
	}
}
func Unsubscribe(username string) {
	nothing := []byte("")
	for j := 1; j <= 6; j++ {
		js := strconv.Itoa(j)
		filename := "name" + js
		find, _ := ioutil.ReadFile(filename)
		if username == string(find) {
			errN := ioutil.WriteFile(filename, nothing, 0644)
			if errN != nil {
				panic(errN)
			}
			passwordFile := "password" + js
			errP := ioutil.WriteFile(passwordFile, nothing, 0644)
			if errP != nil {
				panic(errP)
			}
			break
		}
	}
	database[username] = ""
}
func ClearData() {
	none := ""
	nothing := []byte(none)
	for j := 1; j <= 6; j++ {
		js := strconv.Itoa(j)
		filename := "name" + js
		user, _ := ioutil.ReadFile(filename)
		username := string(user)
		errN := ioutil.WriteFile(filename, nothing, 0644)
		if errN != nil {
			panic(errN)
		}
		passwordFile := "password" + js
		errP := ioutil.WriteFile(passwordFile, nothing, 0644)
		if errP != nil {
			panic(errP)
		}
		database[username] = ""
	}
}
