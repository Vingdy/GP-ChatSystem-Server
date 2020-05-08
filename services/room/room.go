package room

import (
	"GP/db"
	"log"
)

func BanRoom(roomname string) (err error) {
	updateSql := "update gp.room set isban = 1 where roomname = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("BanRoom updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(roomname)
	if err != nil {
		log.Println("BanRoom exec fail")
		return err
	}
	return nil
}

func CancelBanRoom(roomname string) (err error) {
	updateSql := "update gp.room set isban = 0 where roomname = ?"
	stmt, err := db.DB.Prepare(updateSql)
		if err != nil {
		log.Println("CancelBanRoom updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(roomname)
	if err != nil {
		log.Println("CancelBanRoom exec fail")
		return err
	}
	return nil
}

func CreateRoom(roomname string) (err error) {
	insertSql := "insert into gp.room(roomname, isban) values (?, 0)"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		log.Println("CreateRoom insertSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(roomname)
	if err != nil {
		log.Println("CreateRoom exec fail")
		return err
	}
	return nil
}

type RoomInfo struct {
	RoomName string `json:"roomname"`
}

func GetOneRoom(roomname string) (roomInfo []*RoomInfo, err error) {
	roomInfo = []*RoomInfo{}
	querySql := "select roomname from gp.room where isban = 0 and roomname = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetOneRoom Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(roomname)
	if err != nil {
		log.Println("GetOneRoom Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var room RoomInfo
		err := rows.Scan(&room.RoomName)
		if err != nil {
			return nil, err
		}
		roomInfo = append(roomInfo, &room)
	}
	return roomInfo, nil
}

func GetRoomList() (roomInfo []*RoomInfo, err error) {
	roomInfo = []*RoomInfo{}
	querySql := "select roomname from gp.room where isban = 0;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetRoomList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println("GetRoomList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var room RoomInfo
		err := rows.Scan(&room.RoomName)
		if err != nil {
			return nil, err
		}
		roomInfo = append(roomInfo, &room)
	}
	return roomInfo, nil
}


