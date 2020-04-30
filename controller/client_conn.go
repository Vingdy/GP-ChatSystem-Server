package controller

type ClientConn struct {}

/*type ConnInfo struct {
	Token string
	Name string
	Conn *websocket.Conn
}

var ClientConnsMap map[string]ConnInfo

func init() {
	ClientConnsMap = make(map[string]ConnInfo)
}

func (cc ClientConn) Save(name string, conninfo ConnInfo) {
	ClientConnsMap[name] = conninfo
}

func (cc ClientConn) Del(conninfo ConnInfo) {
	for name, connInfo := range ClientConnsMap {
		if conninfo == connInfo {
			delete(ClientConnsMap, name)
		}
	}
}

func (cc ClientConn) SearchByUserName(user model.User) (connInfo *websocket.Conn, err error) {
	//name, err := strconv.Atoi(user.UserName)
	connInfo = ClientConnsMap[user.UserName].Conn
	return
}*/
