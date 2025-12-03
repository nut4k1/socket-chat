package broker

var LastID string = "0-0"

func UpdateLastID(newID string) {
	LastID = newID
}
