package history

import (
	"GP/db"
	"log"
)

type HistoryInfo struct {
	UserName  string `json:"username"`
	Chat      string `json:"chat"`
	Label     string `json:"label"`
	FontType  string `json:"fonttype"`
	FontColor string `json:"fontcolor"`
	Time      string `json:"time"`
}

func NewHistory(roomid, username, chat, label, fonttype, fontcolor, time string) (err error) {
	insertSql := "insert into gp.history(roomname, username, chat, label, fonttype, fontcolor, time) values (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		log.Println("NewHistory insertSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(roomid, username, chat, label, fonttype, fontcolor, time)
	if err != nil {
		log.Println("NewHistory exec fail")
		return err
	}
	return nil
}

func GetHistoryList(roomname string) (historyInfo []*HistoryInfo, err error) {
	historyInfo = []*HistoryInfo{}
	querySql := "select username, chat, label, fonttype, fontcolor, time from gp.history where roomname = ? order by time desc limit 5;"
	stmt, err := db.DB.Prepare(querySql)
	if err != nil {
		log.Println("GetHistoryList Querysql prepare fail")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(roomname)
	if err != nil {
		log.Println("GetHistoryList Querysql query fail")
		return nil, err
	}
	for rows.Next() {
		var history HistoryInfo
		err := rows.Scan(&history.UserName, &history.Chat, &history.Label, &history.FontType, &history.FontColor, &history.Time)
		if err != nil {
			return nil, err
		}
		historyInfo = append(historyInfo, &history)
	}
	return historyInfo, nil
}
