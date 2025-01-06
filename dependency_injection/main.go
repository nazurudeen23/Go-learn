package main

import "fmt"

//dependency injection
type RockClimber struct {
	kind int
	rocksClimbed int
}

func(rc *RockClimber) climbRock(){
	rc.rocksClimbed ++

	if rc.rocksClimbed == 10{
		rc.placeSafties()
	}

}

func (rc *RockClimber) placeSafties(){

	switch rc.kind {
	case 1:
		//ICE
	case 2:
		//Sand
	case 3:
		//concrete
		
	}
	fmt.Println("Placing my safty")
	
}

func main()  {
	rc:=&RockClimber{}

	for i:= 0; i<11;i++{
		rc.climbRock()
	}
}