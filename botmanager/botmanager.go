package botManager

import(
   "os"
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
   "encoding/json"
)

type Headers struct {
   Accept string `json:"accept"`
   Content_Type string `json:"content_type"`
   Authorization string `json:"authorization"`
}

type Msg struct {
   Role string `json:"role"`
   Content string `json:"content"`
}

type Payload struct {
   Model string `json:"model"`
   Maxtoken int `json:"maxtoken"`
   Top_Percent int `json:"top_percent"`
   Top_K int `json:"top_k"`
   Top_N int `json:"top_n"`
   PersencePenality  int `json:"persence_penality"`
   Frequency int `json:"frequency"`
   Temperature int `json:"temperature"`
   Message  []Msg `json:"messages"`
   Stream bool `json:"stream"`
}

type BotManager struct {
   Type         string          `json:"typez"`  // bot 類別
   Url          string          `json:"url"`    // bot Url
   Model        string          `json:"model"`  // bot 模型
   Messages      []Msg           `json:"message"`   
}

// 清空訊息
func(app *BotManager) ClearMessage() {
   app.Messages = []Msg {}
} 

// 設定訊息
func(app *BotManager) SetMessage(message string) {
   app.Messages = append(app.Messages, Msg {
      Role: "user",
      Content: message,
   })
}

// 取得 Payload
func(app *BotManager) GetPayload() (*strings.Reader) {
   pld := Payload {
      Model: app.Model,
      Maxtoken: 4096,
      Top_Percent: 10,
      Top_K: 10,
      Top_N: 10,
      PersencePenality: 0,
      Frequency: 0,
      Temperature: 0,
      Message:app.Messages,
      Stream: false,
   }
   s, err := json.Marshal(pld)
   if err!= nil {
      fmt.Println(err.Error())
      return nil
   }
   return strings.NewReader(string(s))
}

func(app *BotManager) GetResponse() (string, error) {
   client := &http.Client{}
   req, err := http.NewRequest("POST", app.Url, app.GetPayload())
   if err!= nil {
      return "", err
   }
   req.Header.Add("Accept", "application/json")
   req.Header.Add("Content-Type", "application/json")
   req.Header.Add("authorization", "Bearer " + os.Getenv(app.Type + "Token"))
   res, err := client.Do(req)
   if err!= nil {
      return "", err
   }
   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body)
   if err!= nil {
      return "", err
   }
   return string(body), nil
}

// 清空 message
func(app *BotManager) reset(w http.ResponseWriter, r *http.Request) {
   app.ClearMessage()
   fmt.Fprintf(w, "ok")
}

// 詢問 AI
func(app *BotManager) Ask(w http.ResponseWriter, r *http.Request) {
   var str map[string]string
   if r.Method == http.MethodPost {
      if err := r.ParseForm(); err != nil {
         fmt.Fprintf(w, "%v", err)
         return
      }      
     // str["message"] = r.Form.Get("message")
     // if str["message"] == "" {
         b, err := ioutil.ReadAll(r.Body)
         defer r.Body.Close()
         if err != nil {
            fmt.Fprintf(w, "Error: %v, If you use HTMx try post data.", err)
            return
         } else {
            if err := json.Unmarshal(b, &str); err != nil {
               fmt.Fprintf(w, "Error: %v, If you use AngularJS try post data.", err)
               return
            }      
         }
     // }
   } else {      
      webVars := r.URL.Query()
      str["message"] =  webVars.Get("message")
   }
   fmt.Println(str["message"])
   app.SetMessage(str["message"])
   body, err := app.GetResponse()
   if err!= nil {
      fmt.Fprintf(w, "Error: %v, Post data error.", err)
      return
   }
   ans := app.ParseResponse(string(body))
   w.WriteHeader(http.StatusOK)
   w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(ans))   // fmt.Fprintf(w, ans)
   fmt.Println("    " + ans)
}

func(app *BotManager) AddRouter(router *http.ServeMux) {
   router.Handle("/askai", http.HandlerFunc(app.Ask))
   router.Handle("/reset", http.HandlerFunc(app.reset))
}

func NewBotManager(botType string)(*BotManager, error) {
   if botType == "" {
      botType = "taide"   // 預設
   }
   botType = strings.ToLower(botType)
   url := os.Getenv(botType + "Url")
   model := os.Getenv(botType + "Model")
   if model == "" {
      model = "ryan4559/llama3-taide-lx-8b-chat-alpha1-4bit"    // 預設台德
   }
   return &BotManager {
      Type: botType,
      Model: model,
      Url: url,
   }, nil
}