// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(id int) {
// 	fmt.Println("Doing Task", id)
// }

// func main() {

// 	for i := 0; i <= 10; i++ {
// 		// go task(i)
// 	}
// 	time.Sleep(time.Second * 2)

// }

// package main
// import (
// 	"fmt"
// 	"time"
// )
// func task(id int) {
// 	fmt.Println("Doing task", id)
// }
// func main() {
// 	for i := 0; i <= 10; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 2)
// }

//WaitGroup

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func task(id int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	fmt.Println("Doing task", id)
// }

// func main() {
// 	var wG sync.WaitGroup

// 	for i := 0; i <= 10; i++ {
// 		wG.Add(1)
// 		go task(i, &wG)
// 	}
// 	wG.Wait()

// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task Started")
	time.Sleep(time.Second * 2)
	fmt.Println("Task Completed")
}
func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task2 Started")
	time.Sleep(time.Second * 2)
	fmt.Println("Task2 Completed")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go task(&wg)
	go task2(&wg)
	wg.Wait()

	fmt.Println("Main function is Completed")

}

// ```
// int minMoves(int arr_count, int* arr) {
//     int zeros = 0;
//     int ones = 0;

//     int movesZeroFirst = 0;
//     int movesOneFirst = 0;

//     for (int i = 0; i < arr_count; i++) {
//         if (arr[i] == 0) {
//             movesZeroFirst += ones;
//             zeros++;
//         } else {
//             movesOneFirst += zeros;
//             ones++;
//         }
//     }

//     return (movesZeroFirst < movesOneFirst)
//            ? movesZeroFirst
//            : movesOneFirst;
// }
// ```
