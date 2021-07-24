package reportService

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/smtp"
	"ocg-be/database"
	"ocg-be/rmq-sendmail/rabbitmqService"
	"text/template"

	"github.com/streadway/amqp"
)

type Report struct {
	Day string
	Revenue float64
  }
  func getAdminEmail() []string {
  
	  rows, _ := database.DB.Query("SELECT email FROM users u JOIN roles r ON u.role_id = r.id WHERE r.role = 'Admin';")
	  var emails []string
	  for rows.Next() {
		  var email string
		  rows.Scan(&email)
		  emails = append(emails, email)
	  }
	  return emails
  }
  
  func revenueYesterday() Report {
	  rows, _ := database.DB.Query("SELECT create_at, SUM(total_price) FROM orders WHERE create_at = DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY);")
  
	var result Report
	for rows.Next() {
		  rows.Scan(&result.Day,&result.Revenue)
	  }
  
	return result
  }
  
  func covertStructToBytes(input Report) (result []byte){
  
	  reqBodyBytes := new(bytes.Buffer)
	  json.NewEncoder(reqBodyBytes).Encode(input)
	  result = reqBodyBytes.Bytes()
	  return result
  }
  func sendMessage(queue amqp.Queue, ch *amqp.Channel, message Report) {
	  err := ch.Publish(
		  "",
		  queue.Name,
		  false,
		  false,
		  amqp.Publishing{
			  ContentType: "text/plain",
			  Body:        covertStructToBytes(message),
		  })
	  rabbitmqService.FailOnError(err, "Failed to publish message")
  }
  func sendEmailToAdmin(filename Report) {
  
	  // Sender data.
	  from := "hoanganhht1999@gmail.com"
	  password := "alvcgtkytvwtprkd"
  
	  // Receiver email address.
	  to := getAdminEmail()
  
	  // smtp server configuration.
	  smtpHost := "smtp.gmail.com"
	  smtpPort := "587"
  
	  // Authentication.
	  auth := smtp.PlainAuth("", from, password, smtpHost)
  
	  t, _ := template.ParseFiles("index.html")
  
	  var body bytes.Buffer
  
	  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	  body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))
		t.Execute(&body, filename)
  
	  // Sending email.
	  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	  if err != nil {
		  fmt.Println(err)
		  return
	  }
	  fmt.Println("Email Sent!")
  }
  