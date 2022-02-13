package api

type ApiResponse struct {
	Status		uint			`json:"status"`
	Data		interface{}		`json:"data"`
	Message		string			`json:"message"`
}

func (res *ApiResponse) GetStatus() uint {return res.Status}
func (res *ApiResponse) GetData() interface{} {return res.Data}
func (res *ApiResponse) GetMessage() string {return res.Message}

func (res *ApiResponse) SetStatus(status uint) {res.Status = status}
func (res *ApiResponse) SetData(data interface{}) {res.Data = data}
func (res *ApiResponse) SetMessage(message string) {res.Message = message}
