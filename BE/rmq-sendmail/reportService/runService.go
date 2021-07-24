package reportService

import (
	"ocg-be/rmq-sendmail/rabbitmqService"
	"github.com/robfig/cron/v3"
	"time"
)

func RunReportService() {
	conn := rabbitmqService.Connect()
	defer conn.Close()
	ch, _ := conn.Channel()
	defer ch.Close()
	q, err := rabbitmqService.Producer(ch, "order")
	rabbitmqService.FailOnError(err, "can not open queue")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 22h0m0s", func() {
			sendMessage(q, ch, revenueYesterday())
	})

	c.AddFunc("0 17 15 * * * ", func() {
		rabbitmqService.Receiver(ch, q)
		// go func() {}()
		time.Sleep(1 * time.Second)
		sendEmailToAdmin(revenueYesterday())
	})
	c.Start()
	defer c.Stop()
	forever := make(chan bool)
	<-forever
}
