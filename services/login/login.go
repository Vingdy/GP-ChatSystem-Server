package login

import (
	"GP/db"
	"GP/model"
	"database/sql"
	"log"
)

func LoginAccCheck(username string) (ok bool, err error) {
	var haveusername string
	querySql := "select username from gp.user where username = ?"
	err = db.DB.QueryRow(querySql, username).Scan(&haveusername)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("LoginAccCheck not account found")
			return false, nil
		} else {
			log.Println("LoginAccCheck Querysql query fail" + err.Error())
			return true, err
		}
	}
	return true, nil
}

func Login(username, password string) (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select username,nickname,role from gp.user where username = ? and password = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("Login Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, password)
	if err != nil {
		log.Println("Login Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserName, &user.NickName, &user.Role)
		if err != nil {
			return nil, err
		}
		userInfo = append(userInfo, &user)
	}
	return userInfo, nil
}

func GetPassword(username string) (password string, err error) {
	//password = string
	querySql := "select password from gp.user where username = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetPassword Querysql prepare fail")
		return "", err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("GetPassword Querysql query fail")
		return "", err
	}
	for rows.Next() {
		var user string
		err := rows.Scan(&user)
		if err != nil {
			return "", err
		}
		password = user
	}
	return password, nil
}
