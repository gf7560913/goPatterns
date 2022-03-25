package single_mode

import (
	"sync"
)

//##单例模式：
//优点：
//单例模式可以保证内存里只有一个实例，减少了内存的开销。
//可以避免对资源的多重占用。
//单例模式设置全局访问点，可以优化和共享资源的访问。
//缺点：
//单例模式一般没有接口，扩展困难。如果要扩展，则除了修改原来的代码，没有第二种途径，违背开闭原则。
//在并发测试中，单例模式不利于代码调试。在调试过程中，如果单例中的代码没有执行完，也不能模拟生成一个新的对象。
//单例模式的功能代码通常写在一个类中，如果功能设计不合理，则很容易违背单一职责原则。

type Car struct {
	Color string
	Speed string
}
type Ship struct {
	Color string
}
type AirPlane struct {
	Color string
}
type HighTrain struct {
	Color string
}

var (
	Global     *Car
	CarOnce    sync.Once
	GlobalShip *Ship
	Airplane   *AirPlane
	Train      *HighTrain
)

//simple1 线程安全
func NewCar() *Car {
	CarOnce.Do(func() {
		Global = &Car{
			Color: "red",
			Speed: "100km/h",
		}
	})
	return Global
}

//非单例模式
func NewShip() *Ship {
	GlobalShip = &Ship{
		Color: "white",
	}
	return GlobalShip
}

//simple2　线程非安全
func NewAirPlane() *AirPlane {
	if Airplane == nil {
		Airplane = &AirPlane{
			Color: "White",
		}
	}
	return Airplane
}

var lock *sync.Mutex = &sync.Mutex{}

func NewTrain() *HighTrain {
	lock.Lock()
	defer lock.Unlock()
	if Train == nil {
		Train = &HighTrain{
			Color: "White",
		}
	}
	return Train
}
