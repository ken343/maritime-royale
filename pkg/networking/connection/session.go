package connection

var SessionID string

func SetID(ID string) {
	SessionID = ID
}

func GetID() string {
	return SessionID
}
