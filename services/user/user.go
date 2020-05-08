package user

import (
	"GP/db"
	"log"
	"GP/model"
)

func GetUser(username string) (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select username, password, nickname, role, phone, label, head, isban from gp.user where username = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetUser Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("GetUser Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.Role , &user.Phone, &user.Label, &user.Head, &user.IsBan)
		if err != nil {
			return nil, err
		}
		userInfo = append(userInfo, &user)
	}
	return userInfo, nil
}

func UpdateUser(username, nickname, phone, label, head string) (err error) {
	updateSql := "update gp.user set nickname = ?, phone = ?, label = ?, head = ? where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpdateUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nickname, phone, label, head, username)
	if err != nil {
		log.Println("UpdateUser exec fail")
		return err
	}
	return nil
}

func UpdatePassword(username, password string) (err error) {
	updateSql := "update gp.user set password = ? where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpdatePassword updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(password, username)
	if err != nil {
		log.Println("UpdatePassword exec fail")
		return err
	}
	return nil
}

func BanUser(username string) (err error) {
	updateSql := "update gp.user set isban = 1 where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("BanUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username)
	if err != nil {
		log.Println("BanUser exec fail")
		return err
	}
	return nil
}

func CancelBanUser(username string) (err error) {
	updateSql := "update gp.user set isban = 0 where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("CancelBanUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username)
	if err != nil {
		log.Println("CancelBanUser exec fail")
		return err
	}
	return nil
}