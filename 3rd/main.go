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
	for mess!="5"{
		fmt.Scanln(&mess)
		switch mess{
			case "1":
				fmt.Println(b.Hello())
			
			case "2":
				fmt.Println(b.Time())
			
			case "3":
				fmt.Println(b.Date())
			
			case "4":
				fmt.Println(b.Weekday())

			case "5":
				fmt.Println(b.Bye())
				break

			default:
				fmt.Println("There are not such command. Try again.")
		}
	}

}