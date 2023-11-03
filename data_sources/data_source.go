package data_sources

type DataSource interface {
	PushData(data string) error
}
