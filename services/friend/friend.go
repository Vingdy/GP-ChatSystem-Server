package friend

import (
	"GP/db"
	"log"
	"database/sql"
)

type FriendInfo struct {
	Id         string `json:"id"`
	FriendName string `json:"friendname"`
}

func GetCheckFriend(username string) (friendInfo []*FriendInfo, err error) {
	friendInfo = []*FriendInfo{}
	querySql := "select id, nickname1 from gp.friend where username2 = ? and ischeck = 0;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetCheckFriend Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("GetCheckFriend Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var friend FriendInfo
		err := rows.Scan(&friend.Id, &friend.FriendName)
		if err != nil {
			return nil, err
		}
		friendInfo = append(friendInfo, &friend)
	}
	return friendInfo, nil
}

func GetFriendList(username string) (friendInfo []*FriendInfo, err error) {
	friendInfo = []*FriendInfo{}
	querySql := "select id, nickname2 from gp.friend where username1 = ? and ischeck = 1;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetFriendList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println("GetFriendList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var friend FriendInfo
		err := rows.Scan(&friend.Id, &friend.FriendName)
		if err != nil {
			return nil, err
		}
		friendInfo = append(friendInfo, &friend)
	}
	return friendInfo, nil
}

func NewFriend(username1, nickname1, username2, nickname2 string) (err error) {
	insertSql := "insert into gp.friend(username1, nickname1, username2, nickname2, ischeck) values (?, ?, ?, ?, 0)"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
	log.Println("NewFriend insertSql fail")
	return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username1, nickname1, username2, nickname2)
	if err != nil {
	log.Println("NewFriend exec fail")
	return err
	}
	return nil
}

func PassFriendIdCheck(username string) (ok bool, err error) {
	var haveid string
	querySql := "select id from gp.friend where id = ?"
	err = db.DB.QueryRow(querySql, username).Scan(&haveid)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("PassFriendIdCheck not account found")
			return false, nil
		} else {
			log.Println("PassFriendIdCheckQuerysql query fail" + err.Error())
			return true, err
		}
	}
	return true, nil
}

func PassFriend(id string) (err error) {
	updateSql := "update gp.firend set ischeck = 1 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("PassFriend updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("PassFriend exec fail")
		return err
	}
	return nil
}

