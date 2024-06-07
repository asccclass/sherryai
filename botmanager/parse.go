package botManager

import(
   // "fmt"
   "encoding/json"
)

type TAIDEresponse struct {
	Model		string	`json:"model"`
	CreateedAt		string	`json:"created_at"`
	Message		Msg	`json:"message"`
	DoneReason		string	`json:"done_reason"`
	Done			bool	`json:"done"`
	TotalDuration	int	`json:"total_duration"`
	LoadDuration	int	`json:"load_duration"`
	PromptEvalCount	int	`json:"prompt_eval_count"`
	PromptEvalDuration	int	`json:"prompt_eval_duration"`
	EvalCount	int	`json:"eval_count"`
	EvalDuration	int	`json:"eval_duration"`
}

func(app *BotManager) GetAIDEResponse(response string)(string) {
	var responseTAIDE TAIDEresponse
	json.Unmarshal([]byte(response), &responseTAIDE)
	return responseTAIDE.Message.Content
}

func(app *BotManager) ParseResponse(response string) (string) {
	str := ""
   if app.Type == "taide" {
      str = app.GetAIDEResponse(response)
   } else {
		return response
	}
	s := make(map[string]string) // map[string]string
	s["ans"] = str
	jsondata, _ := json.Marshal(s)
	return string(jsondata)
}