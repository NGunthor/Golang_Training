package pattern_proxy

// DBFunction is the common interface for data-base, db-node and cache
type DataBaseController interface {
	GetData() string
	SendData(data string)
}
