package gitlink

// 存放公共结构体

type Author struct {
	Id       int32  `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ImageUrl string `json:"image_url"`
}
