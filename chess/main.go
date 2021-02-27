
package chess

import(
	os    "os"
)

func RunWindow(){
	INFO.Println("Welcome to use this AI chess player")
	INFO.Printf("The screen width is %d, height is %d\n", SCREEN_WIDTH, SCREEN_HEIGHT)

	_, err := GetWindow()
	if err != nil{
		os.Exit(128)
	}
	defer CloseWindow()
	for !WindowShouldClose(){
		DrawWindow()
		PollEvents()
	}

	INFO.Println("Thank you for use this AI chess player")
	INFO.Println("Byby")
}

func RunText(){
	INFO.Println("Welcome to use this AI chess player")
	INFO.Printf("The screen width is %d, height is %d\n", SCREEN_WIDTH, SCREEN_HEIGHT)

	//

	INFO.Println("Thank you for use this AI chess player")
	INFO.Println("Byby")
}