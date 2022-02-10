package store

type AuthStore interface {
	Save(entry ClusterInfo) error
	Get() ClusterInfo
}

type ClusterInfo struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
	Current  bool   `json:"current"`
}
type LocalFileAuthStore struct {
}

func (l LocalFileAuthStore) Save(entry ClusterInfo) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalFileAuthStore) Get() ClusterInfo {
	//TODO implement me
	panic("implement me")
}
