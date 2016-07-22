package timermanager

import (
	"github.com/RexGene/common/timingwheel"
)

type TimerManager struct {
	data map[uint]*timingwheel.TimingWheel
}

var instance *TimerManager
var hmsInstance *TimerManager

func New() *TimerManager {
	return &TimerManager{
		data: make(map[uint]*timingwheel.TimingWheel),
	}
}

func GetInstance() *TimerManager {
	if instance == nil {
		instance = New()
	}

	return instance
}

func GetHmsInstance() *TimerManager {
	if hmsInstance == nil {
		hmsInstance = New()
	}

	return hmsInstance
}

func (self *TimerManager) Tick() {
	for _, tw := range self.data {
		tw.Tick()
	}
}

func (self *TimerManager) AddTimer(sec uint, callback func()) *timingwheel.BaseNode {
	tw := self.data[sec]
	if tw == nil {
		tw = timingwheel.New(sec)
		self.data[sec] = tw
	}

	return tw.InsertCallback(callback)
}

func (self *TimerManager) AddTimerForever(sec uint, callback func()) *timingwheel.BaseNode {
	tw := self.data[sec]
	if tw == nil {
		tw = timingwheel.New(sec)
		self.data[sec] = tw
	}

	return tw.InsertCallbackForever(callback)
}

func (self *TimerManager) RemoveTimer(base *timingwheel.BaseNode) {
	timingwheel.Remove(base)
}

func RemoveTimer(base *timingwheel.BaseNode) {
	timingwheel.Remove(base)
}
