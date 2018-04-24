package docutils

//import "encoding/json"

type MessageHeader struct {
	Name string
	Version string
}

type ProgressMessage struct {
	MeetingId string
	Room string
	ReturnCode string
	PresentationName string
	PresentationId string
	Filename string
}

type SlideGeneratedMessage struct {
	MeetingId string
	MessageKey string
	Code string
	PresentationId string
	NumPages string
	PagesCompleted string
	PresentationName string
}

type OfficeConversionSuccessMessage struct {
	MeetingId 			string 		`json:"conference"`
	Room				string		`json:"room"`
	ReturnCode			string		`json:"returnCode"`
	PresentationId		string		`json:"presentationId"`
	PresentationName 	string		`json:"presentationName"`
	Filename			string		`json:"filename"`
}