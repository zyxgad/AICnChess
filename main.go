
package main

import(
	os    "os"

	chess "./github.com/zyxgad/AICnChess/chess"
)


func init(){
}

func main(){
	chess.INFO.Println("Welcome to use this AI chess player")
	chess.INFO.Printf("The screen width is %d, height is %d\n", chess.SCREEN_WIDTH, chess.SCREEN_HEIGHT)

	_, err := chess.GetWindow()
	if err != nil{
		os.Exit(128)
	}
	defer chess.CloseWindow()
	for !chess.WindowShouldClose(){
		chess.DrawWindow()
		chess.PollEvents()
	}

	chess.INFO.Println("Thank you for use this AI chess player")
	chess.INFO.Println("Byby")
}