package entity

type Task struct {
	Id             string
	Method         string
	Url            string
	Headers        map[string]string
	ReqBody        string
	RespBody       string
	Status         string
	HttpStatusCode int
	Length         int
}
