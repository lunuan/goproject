package consul

type ConsulClient struct {
	Address string
	KV      ConsulKVData
}

type ConsulKVData struct {
}

func (ckv ConsulKVData) Get(key string) interface{} {
	return "None"
}
