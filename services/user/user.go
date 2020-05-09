package user

import (
	"GP/db"
	"log"
	"GP/model"
	"database/sql"
)

func GetOneUser(username string) (userInfo []*model.User, err error) {
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

func UpUserRole(username string) (err error) {
	updateSql := "update gp.user set role = 'manager' where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpUserRole updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username)
	if err != nil {
		log.Println("UpUserRole exec fail")
		return err
	}
	return nil
}

func DownUserRole(username string) (err error) {
	updateSql := "update gp.user set role = 'member' where username = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("DownUserRoleupdateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username)
	if err != nil {
		log.Println("DownUserRole exec fail")
		return err
	}
	return nil
}

func FindUser(findstring string) (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select username, password, nickname, role, phone, label, head, isban from gp.user where username like ? or nickname like ? or label like ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("FindUser Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+findstring+"%", "%"+findstring+"%", "%"+findstring+"%")
	if err != nil {
		log.Println("FindUser Querysql query fail")
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

func GetUserRole(username string) (role string, err error) {
	querySql := "select role from gp.user where username = ?"
	err = db.DB.QueryRow(querySql, username).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user not account found")
			return "null", nil
		} else {
			log.Println("GetUserRole query fail" + err.Error())
			return "null", err
		}
	}
	return role, nil
}

