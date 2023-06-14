package v1

import (
	"github.com/wsmbsbbz/predictor/util"
	"reflect"
)

type Interface interface {
	Drawer
	GoalChecker
}

type Drawer interface {
	// Draw returns the result of a single draw(in the form of type int).
	// note: 完成整个抽奖流程，如果有粉球数量需求，应该在GenshinDrawer中实现
	Draw()
}

type GoalChecker interface {
	CheckAchieve() bool
}

// func (p Predictor) Result() string {
// 	return fmt.Sprintf("The predictor has simulated %d times, %d of which were successful. The successful possibility is %.2f",
// 		p.PredictTimes, p.AchieveTimes, float64(p.PredictTimes)/float64(p.AchieveTimes)*100)
// }

// Predict run n times of predicts, return the successful times.
func Predict(p Interface, n int) int {
	successfulTimes := 0
	for i := 1; i <= n; i++ {
		// cur := p.Copy().(Interface)
		cur := deepCopyInterface(p).(Interface)
		cur.Draw()
		if cur.CheckAchieve() {
			successfulTimes++
		}
	}
	util.Dprintf("Predict: %d/%d\n", successfulTimes, n)
	return successfulTimes
}

// Copeid from ChatGPT
// 深度复制接口的值
func deepCopyInterface(i interface{}) interface{} {
	// 如果传入的是nil，则直接返回nil
	if i == nil {
		return nil
	}

	// 获取接口的反射类型和值
	originalType := reflect.TypeOf(i)
	originalValue := reflect.ValueOf(i)

	// 如果传入的是指针类型，则复制指针所指向的对象
	if originalType.Kind() == reflect.Ptr {
		originalValue = originalValue.Elem()
	}

	// 创建一个新的对象并复制值
	cloneValue := reflect.New(originalValue.Type()).Elem()
	cloneValue.Set(originalValue)

	// 如果原始对象是指针类型，则返回复制后的指针值
	if originalType.Kind() == reflect.Ptr {
		clonePtr := cloneValue.Addr().Interface()
		return clonePtr
	}

	// 返回复制后的非指针值
	clone := cloneValue.Interface()
	return clone
}
