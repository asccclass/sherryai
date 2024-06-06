package botManager

import(
   "os"
   // "fmt"
)

type BotManager struct {
   Type         string          `json:"typez"`  // bot 類別
   Url          string          `json:"url"`    // bot Url
}

func NewBotManager(botType string)(*BotManager, error) {
   if botType == "" {
      botType = "Taide"   // 預設
   }
   url := os.Getenv(botType + "Url")

   return &BotManager {
      Type: botType,
      Url: url,
   }, nil
}