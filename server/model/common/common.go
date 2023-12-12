package common

// CustomError error contains error message and status code
type CustomError struct {
	Message    error
	StatusCode int
}


// type CustomMessage struct{
// 	MessageErr error
// 	Message    string
// 	StatusCode  int
// 	TypeResult  string
// }

// DownloadReport download reports struct
type DownloadReport struct {
	FileName    string `json:"fileName" form:"fileName"`
	DownloadUrl string `json:"downloadUrl" form:"downloadUrl"`
}


type GData struct{
	Key string 
	Value string
	// FileID string
}