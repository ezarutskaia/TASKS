package main

import (
    "log"
    "strconv"
    "strings"
    "tasks/utils"
    "tasks/models"
    "tasks/rabbit"
    "encoding/json"
    "github.com/jung-kurt/gofpdf"
)


func main() {
    db := *utils.Engine()
    ch, _ := rabbit.ConnectRabbitMQ()
    var tasks []models.Task
    var ids []string

    q, err := ch.QueueDeclare(
		"id_task",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

    msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

    for d := range msgs {
        err := json.Unmarshal(d.Body, &ids)
        if err != nil {
            log.Printf("Failed to unmarshal message: %s", err)
            continue
        }

        result := db.Where("id IN ?", ids).Find(&tasks)
        if result.Error != nil {
            log.Fatalf("Failed to find users: %v", result.Error)
        }

        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 16)
        pdf.Cell(40, 10, "Task List")
        pdf.Ln(12)

        pdf.SetFont("Arial", "", 12)
        for _, task := range tasks {
            pdf.Cell(40, 10, strconv.Itoa(task.Id))
            pdf.Cell(60, 10, task.Name)
            pdf.Ln(10)
        }

        err = pdf.OutputFileAndClose("pdf/tasks-" + strings.Join(ids, ", ") + ".pdf")
        if err != nil {
            log.Fatalf("Failed to generate PDF: %v", err)
        }

        log.Println("PDF successfully generated!")
    }
}
