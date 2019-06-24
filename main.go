package main
import (
	"fmt"
)



func main() {
	var pogramnum int32
	fmt.Println("Hello Paramananda . I'm in my Day-2 Test")
	fmt.Println("Enter to Quit.  1: Shuffle . 2:calculator , 3:Count Days , 4:ArraysInGolang , 5:Email ")
	fmt.Scanf("%d",&pogramnum)
	switch pogramnum {
	case 1:
		Shuffle()
	case 2:
		calculator()
	case 3:
		countDays()
	case 4:
		ArraysInGolang()
	case 5:
		eMailInGolang()
	default:
		fmt.Println("Bye. Have a good day.")

	}

}
