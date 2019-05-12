package main
import(
	"fmt"
	"3rd/bots"
	"errors"
	"os"
)

func main(){
	bot, err:=createBot("Russian")
	if err!=nil{
		fmt.Println("Error has happend: ",err)
		os.Exit(1)
	}
	listenUser(bot)
}

func createBot(language string) (bots.Bot, error){
	switch language{
		case "English":
			var bot bots.Bot=bots.EnglishBot{Name: "John-bot"}
			return bot,nil
		
		case "Russian":
			var bot bots.Bot=bots.RussianBot{Name: "Жора-бот"}
			return bot,nil

		default:
			return nil, errors.New("Unknown laguage")
	}
}
func listenUser(b bots.Bot){
	var mess string
	for {
		fmt.Scanln(&mess)
		fmt.Println(b.Command(mess))

		if mess==b.ExitWord(){
			break
		}
	}
}