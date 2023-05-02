package entity

type Task struct {
	Id             string
	Method         string
	Url            string
	Headers        map[string]string
	ReqBody        []byte
	RespBody       []byte
	Status         string
	HttpStatusCode int
	Length         int
}
