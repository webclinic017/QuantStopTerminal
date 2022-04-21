package trader

import (
	"github.com/d5/tengo/v2"
	"github.com/quantstop/quantstopterminal/internal/log"
	"time"
)

var script = tengo.NewScript([]byte(
	`each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
    sum += x
    mul *= x
})`))

func Run() {
	log.Debugln(log.TraderLogger, "starting workers ...")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		log.Debugf(log.TraderLogger, "worker %v started job %v", id, j)
		time.Sleep(time.Second)

		//fmt.Println("worker", id, "finished job", j)
		log.Debugf(log.TraderLogger, "worker %v finished job %v", id, j)
		results <- j * 2
	}
}
