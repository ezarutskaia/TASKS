package main

import (
    "log"
    "strconv"
    "tasks/utils"
    "tasks/models"
    "github.com/jung-kurt/gofpdf"
)


func GetPdfByID(id []int) {
    db := *utils.Engine()
    var tasks []models.Task

    result := db.Where("id IN ?", id).Find(&tasks)
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

    err := pdf.OutputFileAndClose("pdf/tasks.pdf")
    if err != nil {
        log.Fatalf("Failed to generate PDF: %v", err)
    }

    log.Println("PDF successfully generated!")
}

func main() {

    setID := []int{1, 2, 3, 4, 5}
    GetPdfByID(setID)

}
