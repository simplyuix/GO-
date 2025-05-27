package main

import (
	"fmt"
)

func main()  {


	//  defer is more like async await { runs in backdround compute then come back until the I/O is not free }
        
	  i := 0
fmt.Println("Before Defer:", i)
i++

defer fmt.Println("Deferred!", i)
i++
fmt.Println("After Defer 1:", i)
i++
fmt.Println("After Defer 2:", i)
i++

defer fmt.Println("checkPoint", i)
i++
fmt.Println("After Defer 3:", i)
i++


}