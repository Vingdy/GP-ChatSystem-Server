package user

import (
	"GP/db"
	"log"
	"GP/model"
	"database/sql"
)

func GetOneUser(id string) (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select id, username, nickname, role, phone, label, fonttype, fontcolor, isban from gp.user where id = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetUser Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("GetUser Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.Role , &user.Phone, &user.Label, &user.FontType, &user.FontColor, &user.IsBan)
		if err != nil {
			return nil, err
		}
		userInfo = append(userInfo, &user)
	}
	//fmt.Println(userInfo)
	return userInfo, nil
}

func GetUserList() (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select id, username, nickname, role, phone, label, fonttype, fontcolor, isban from gp.user"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetUserList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println("GetUserList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.Role , &user.Phone, &user.Label, &user.FontType, &user.FontColor, &user.IsBan)
		if err != nil {
			return nil, err
		}
		userInfo = append(userInfo, &user)
	}
	//fmt.Println(userInfo)
	return userInfo, nil
}

func UpdateUser(id, nickname, phone, label, fonttype, fontcolor string) (err error) {
	updateSql := "update gp.user set nickname = ?, phone = ?, label = ?, fonttype = ?,fontcolor = ? where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpdateUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(nickname, phone, label, fonttype, fontcolor, id)
	if err != nil {
		log.Println("UpdateUser exec fail")
		return err
	}
	return nil
}

func UpdatePassword(id, password string) (err error) {
	updateSql := "update gp.user set password = ? where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpdatePassword updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(password, id)
	if err != nil {
		log.Println("UpdatePassword exec fail")
		return err
	}
	return nil
}

func BanUser(id string) (err error) {
	updateSql := "update gp.user set isban = 1 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("BanUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("BanUser exec fail")
		return err
	}
	return nil
}

func CancelBanUser(id string) (err error) {
	updateSql := "update gp.user set isban = 0 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("CancelBanUser updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("CancelBanUser exec fail")
		return err
	}
	return nil
}

func UpUserRole(id string) (err error) {
	updateSql := "update gp.user set role = 'manager' where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UpUserRole updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("UpUserRole exec fail")
		return err
	}
	return nil
}

func DownUserRole(id string) (err error) {
	updateSql := "update gp.user set role = 'member' where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("DownUserRoleupdateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("DownUserRole exec fail")
		return err
	}
	return nil
}

func FindUser(findstring string) (userInfo []*model.User, err error) {
	userInfo = []*model.User{}
	querySql := "select id, username, nickname, role, phone, label, fonttype, fontcolor, isban from gp.user where nickname like ? or label like ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("FindUser Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%"+findstring+"%", "%"+findstring+"%")
	if err != nil {
		log.Println("FindUser Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.Role , &user.Phone, &user.Label, &user.FontType, &user.FontColor, &user.IsBan)
		if err != nil {
			return nil, err
		}
		userInfo = append(userInfo, &user)
	}
	return userInfo, nil
}

func GetUserRole(id string) (role string, err error) {
	querySql := "select role from gp.user where id = ?"
	err = db.DB.QueryRow(querySql, id).Scan(&role)
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

