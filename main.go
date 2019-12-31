package main

import (
  "os"
  "fmt"
  "strings"
  "time"

  "./cupfmc"
  "github.com/stianeikeland/go-rpio"
)

func main() {
  if len(os.Args) < 2 {
      fmt.Printf("%s: Usage: %s TEXT...\n", os.Args[0], os.Args[0])
      return
  }

  fmt.Printf("Loading...")

  text := strings.Join(os.Args[1:], " ")
  delays := cupfmc.GetDelays(text)
  SendTX(delays, 17)
}

func SendTX(delays []time.Duration, pinout int) {
  err := rpio.Open()
  if err != nil {
    fmt.Printf("%s: Oops: %s\n", err.Error())
    return
  }

  pin := rpio.Pin(pinout)
  pin.Output()
  length := len(delays)-1
  for i := 0; i < len(delays); i++ {
    fmt.Printf("\rSending TX: %d/%d", i, length)
    pin.Toggle()
    time.Sleep(delays[i])
  }

  pin.Low()
  rpio.Close()

  fmt.Printf("\nTX completed.\n")
}
