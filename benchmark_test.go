package tartara

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkInsert(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {

		var fruit = &Fruit{
			Name:  "Fruit" + string(i),
			Color: "#" + string(i),
		}

		_ = fruit

		Fruits.Insert(fruit)
	}
	end := time.Now()

	fmt.Println("N:", b.N, "\t", "Time (ms):", (end.Nanosecond()-start.Nanosecond())/1000000)
}
