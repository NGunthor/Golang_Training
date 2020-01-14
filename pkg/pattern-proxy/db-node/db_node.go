package db_node

type dataBaseNodeController interface {
	GetData() string
	SendData(data string)
}

type dbNode struct {
	cache    dataBaseNodeController
	dataBase dataBaseNodeController
}

// SendData updates data in the main data-base and cache
func (n *dbNode) SendData(data string) {
	dataBaseNodeController.SendData(n.cache, data)
	dataBaseNodeController.SendData(n.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (n dbNode) GetData() string {
	freshData := dataBaseNodeController.GetData(n.cache)
	return freshData
}

// NewDataBaseNodeController returns new instance of proxy
func NewDataBaseNodeController(cache dataBaseNodeController, db dataBaseNodeController) dataBaseNodeController {
	return &dbNode{cache, db}
}
