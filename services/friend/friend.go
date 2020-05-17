package friend

import (
	"GP/db"
	"log"
	"database/sql"
)

type CheckFriendInfo struct {
	Id string `json:"id"`
	FriendId         string `json:"friendid"`
	UserName string `json:"username"`
	FriendName string `json:"friendname"`
	Label string `json:"label"`
}

func GetCheckFriend(username string) (friendInfo []*CheckFriendInfo, err error) {
	friendInfo = []*CheckFriendInfo{}
	querySql := "select id, id2, username2, nickname2, label2 from gp.friend where username1 = ? and ischeck = 0;"
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
		var friend CheckFriendInfo
		err := rows.Scan(&friend.Id, &friend.FriendId, &friend.UserName, &friend.FriendName, &friend.Label)
		if err != nil {
			return nil, err
		}
		friendInfo = append(friendInfo, &friend)
	}
	return friendInfo, nil
}

type FriendInfo struct {
	Id             string `json:"id"`
	FriendId       string `json:"friendid"`
	UserName       string `json:"username"`
	FriendUserName string `json:"friendusername"`
	FriendName     string `json:"friendname"`
	Label          string `json:"label"`
}

func GetFriendList(username string) (friendInfo []*FriendInfo, err error) {
	friendInfo = []*FriendInfo{}
	querySql := "select id, username1, id2, username2, nickname2, label2 from gp.friend where (username1 = ? or username2 = ?) and ischeck = 1;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetFriendList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, username)
	if err != nil {
		log.Println("GetFriendList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var friend FriendInfo
		err := rows.Scan(&friend.Id, &friend.UserName, &friend.FriendId,  &friend.FriendUserName, &friend.FriendName, &friend.Label)
		if err != nil {
			return nil, err
		}
		friendInfo = append(friendInfo, &friend)
	}
	return friendInfo, nil
}

func NewFriendCheck(username1, username2 string) (ok bool, err error) {
	var have string
	querySql := "select id from gp.friend where username1 = ? and username2 = ? and ischeck = 0"
	err = db.DB.QueryRow(querySql, username1, username2).Scan(&have)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			log.Println("NewFriendCheck Querysql query fail" + err.Error())
			return true, err
		}
	}
	log.Println("NewFriendCheck found")
	return true, nil
}

func NewFriend(username1, nickname1, id2, username2, nickname2, label2 string) (err error) {
	insertSql := "insert into gp.friend(username1, nickname1, id2, username2, nickname2, label2, ischeck) values (?, ?, ?, ?, ?, ?, 0)"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
	log.Println("NewFriend insertSql fail")
	return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username1, nickname1, id2, username2, nickname2, label2)
	if err != nil {
	log.Println("NewFriend exec fail")
	return err
	}
	return nil
}

func PassFriendIdCheck(id string) (ok bool, err error) {
	var haveid string
	querySql := "select id2 from gp.friend where id = ?"
	err = db.DB.QueryRow(querySql, id).Scan(&haveid)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("PassFriendIdCheck not found")
			return false, nil
		} else {
			log.Println("PassFriendIdCheckQuerysql query fail" + err.Error())
			return true, err
		}
	}
	return true, nil
}

func PassFriend(id string) (err error) {
	updateSql := "update gp.friend set ischeck = 1 where id = ?"
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

func UnPassFriend(id string) (err error) {
	updateSql := "update gp.friend set ischeck = -1 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("UnPassFriend updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("UnPassFriend exec fail")
		return err
	}
	return nil
}

