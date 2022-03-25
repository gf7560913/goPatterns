package single_mode

import (
	"fmt"
	"testing"
)

func TestSimple1(T *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("car point is %p \n", NewCar())
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("ship point is %p \n", NewShip())
	}

}
func TestSimple2(T *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("airplane point is %p \n", NewAirPlane())
	}

}
func TestSimple3(T *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("airplane point is %p \n", NewTrain())
	}

}
