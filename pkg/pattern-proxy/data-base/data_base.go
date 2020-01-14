package data_base

type dataBaseController interface {
	GetData() string
	SendData(data string)
}

type dataBase struct {
	data  string
	cache dataBaseController
}

// SendData updates data in the main data-base and cache
func (db *dataBase) SendData(data string) {
	db.data = data
	dataBaseController.SendData(db.cache, data)
}

// GetData returns data from main data-base
func (db dataBase) GetData() string {
	return db.data
}

// NewDataBaseController returns new instance of main data-base
func NewDataBaseController(cache dataBaseController) dataBaseController {
	return &dataBase{cache: cache}
}
