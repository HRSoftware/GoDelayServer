package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"time"
)

//Default delay in milliseconds
var responseDelay_ms int = 1000

func PutDelayMs(context *gin.Context) {
	delay := context.Query("delayMs")

	delayDuration, err := strconv.Atoi(delay)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"Setting delay to %d ms\n",
		delayDuration,
	)

	responseDelay_ms = delayDuration
}

func runDelay(timeMs int, wg *sync.WaitGroup) {
	defer fmt.Printf("Finished delay of %d ms\n", timeMs)
	defer wg.Done()

	time.Sleep(time.Millisecond * time.Duration(timeMs))
}

func handleDelayReq(context *gin.Context) {

	var wg sync.WaitGroup
	var delayDuration int
	var err error

	delay := context.Query("delayMs")

	// If we don't pass a param value
	if delay == "" {
		fmt.Printf(
			"No delay specified - using set delay of %d ms\n\n",
			responseDelay_ms,
		)

		delayDuration = responseDelay_ms
	} else {
		delayDuration, err = strconv.Atoi(delay)
		if err != nil {
			fmt.Printf(
				"Failed to parse param - using set delay of %d ms\n%s\n",
				responseDelay_ms,
				err.Error(),
			)
			delayDuration = responseDelay_ms
		}
	}

	wg.Add(1)
	go runDelay(delayDuration, &wg)

	wg.Wait()
}

func main() {
	router := gin.Default()
	router.GET("/delay", handleDelayReq)
	router.PUT("/delay", PutDelayMs)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

}
