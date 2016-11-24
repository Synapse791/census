package commands

import (
  "os"
  "time"
  "net/http"
  "github.com/olekukonko/tablewriter"
)

func newTableWriter() *tablewriter.Table {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetAlignment(tablewriter.ALIGN_CENTER)
  table.SetBorder(false)
  table.SetCenterSeparator("-")

  return table
}

func newHttpClient() *http.Client {
  return &http.Client{
    Timeout: time.Second * 10,
  }
}
