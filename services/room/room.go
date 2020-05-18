package room

import (
	"GP/db"
	"log"
)

func BanRoom(id string) (err error) {
	updateSql := "update gp.room set isban = 1 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("BanRoom updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("BanRoom exec fail")
		return err
	}
	return nil
}

func CancelBanRoom(id string) (err error) {
	updateSql := "update gp.room set isban = 0 where id = ?"
	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		log.Println("CancelBanRoom updateSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
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

type RoomUseInfo struct {
	Id       string `json:"id"`
	RoomName string `json:"roomname"`
}

type RoomInfo struct {
	Id       string `json:"id"`
	RoomName string `json:"roomname"`
	IsBan    string `json:"isban"`
}

func GetOneRoom(id string) (roomInfo []*RoomUseInfo, err error) {
	roomInfo = []*RoomUseInfo{}
	querySql := "select id, roomname from gp.room where isban = 0 and id = ?;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetOneRoom Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("GetOneRoom Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var room RoomUseInfo
		err := rows.Scan(&room.Id, &room.RoomName)
		if err != nil {
			return nil, err
		}
		roomInfo = append(roomInfo, &room)
	}
	return roomInfo, nil
}

func GetRoomList() (roomInfo []*RoomInfo, err error) {
	roomInfo = []*RoomInfo{}
	querySql := "select id, roomname, isban from gp.room;"
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
		err := rows.Scan(&room.Id, &room.RoomName, &room.IsBan)
		if err != nil {
			return nil, err
		}
		roomInfo = append(roomInfo, &room)
	}
	return roomInfo, nil
}

func GetUseRoomList() (roomInfo []*RoomUseInfo, err error) {
	roomInfo = []*RoomUseInfo{}
	querySql := "select id, roomname from gp.room where isban = 0;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetUseRoomList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println("GetUseRoomList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var room RoomUseInfo
		err := rows.Scan(&room.Id, &room.RoomName)
		if err != nil {
			return nil, err
		}
		roomInfo = append(roomInfo, &room)
	}
	return roomInfo, nil
}
