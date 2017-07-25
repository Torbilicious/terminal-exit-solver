package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

var (
	levels map[int]func() time.Duration
)

func main()  {
	fmt.Println("Terminal exit solver!")

	initLevels()

	focusBrowser()

	solveLevels(1, 7)
}

func initLevels() {
	levels = make(map[int]func() time.Duration)

	levels[1] = solveFirstLevel
	levels[2] = solveSecondLevel
	levels[3] = solveThirdLevel
	levels[4] = solveFourthLevel
	levels[5] = solveFifthLevel
	levels[6] = solveSixthLevel
	levels[7] = solveSeventhLevel
}

func focusBrowser() {
	fmt.Println("focusing browser window")

	//bitmap := robotgo.CaptureScreen(200, 200, 100, 100)
	//fmt.Println("...", bitmap)
	//
	//fx, fy := robotgo.FindBitmap(bitmap)
	//fmt.Println("FindBitmap------", fx, fy)
	//
	//robotgo.SaveBitmap(bitmap, "test.png")

	time.Sleep(time.Second * 2)
}

func solveLevels(from int, to int) {
	var waitTime time.Duration

	for i := from; i <= to; i++ {
		fmt.Println("solving level ", i)

		waitTime = levels[i]()

		time.Sleep(waitTime)
	}
}

func do(command string) {
	robotgo.TypeString(command)
	robotgo.KeyTap("enter")
}

func solveFirstLevel() time.Duration {

	do("exit")
	do("y")

	return time.Second * 8
}

func solveSecondLevel() time.Duration {
	do("update")

	return time.Second * 5
}

func solveThirdLevel() time.Duration {

	do("install exit")

	do("exit")
	do("y")

	return time.Second * 8
}

func solveFourthLevel() time.Duration {

	do("install rm")

	do("rm trash_C2H5OH")

	do("exit")

	return time.Second * 8
}

func solveFifthLevel() time.Duration {

	do("exit")
	do("2517")

	return time.Second * 8
}

func solveSixthLevel() time.Duration {

	do("exit")
	do("4385")

	return time.Second * 8
}

func solveSeventhLevel() time.Duration {

	do("install login")

	do("login")
	do("root")
	do("toor1234")

	do("install exit")
	do("exit")

	return time.Second * 2
}
