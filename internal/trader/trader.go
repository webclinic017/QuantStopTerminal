package trader

import (
	"database/sql"
	"github.com/d5/tengo/v2"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/websocket"
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

func Run(db *sql.DB, hub *websocket.Hub) {

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
