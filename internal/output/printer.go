package output

import (
        "fmt"

        "github.com/tools-tensorsecurityacademy/dirrumble/internal/models"
)

func PrintBanner() {
        fmt.Println(`
   _____          _          _____ _     _       
  / ____|        | |        |  __ (_)   | |      
 | |     ___   __| | ___  __| |__) | |__ | | ___  
 | |    / _ \ / _` + "`" + ` |/ _ \/ _` + "`" + ` |  ___/| '_ \| |/ _ \ 
 | |___| (_) | (_| |  __/ (_| | |    | |_) | |  __/
  \_____\___/ \__,_|\___|\__,_|_|    |_.__/|_|\___|

DirRumble v0.1 — Tensor Security Academy (Team Alpha)
Raw Fuzzing Unleashed
`)
}

func PrintResult(res *models.Result) {
        if res.Error != nil {
                fmt.Printf("[ERR]  %s → %v\n", res.URL, res.Error)
                return
        }

        color := getStatusColor(res.Status)
        timeMs := res.Time.Milliseconds()

        fmt.Printf("[\033[%sm%3d\033[0m] %6d bytes | %4d w | %4d ms | %s\n",
                color, res.Status, res.Length, res.Words, timeMs, res.URL)
}

func getStatusColor(status int) string {
        switch {
        case status >= 200 && status < 300:
                return "32" // green
        case status >= 300 && status < 400:
                return "36" // cyan
        case status >= 400 && status < 500:
                return "33" // yellow
        case status >= 500:
                return "31" // red
        default:
                return "37" // white
        }
}
