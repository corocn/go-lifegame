package main

import (
  "os"
  "os/exec"
  "time"
)

var tate = 6
var yoko = 6
var cells [][]int

func main() {
  cells = [][]int{
    {0, 0, 0, 0, 0, 0},
    {0, 1, 1, 0, 0, 0},
    {0, 1, 1, 0, 0, 0},
    {0, 0, 0, 1, 1, 0},
    {0, 0, 0, 1, 1, 0},
    {0, 0, 0, 0, 0, 0},
  }

  for {
    display()
    cells = nextCells()
    time.Sleep(1 * time.Second)
  }
}

func nextCells() [][]int {
  next_cells := [][]int{
    {0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0},
  }

  for y := 0; y < tate; y++ {
    for x := 0; x < yoko; x++ {
      if isNextAlive(x, y) == true {
        next_cells[y][x] = 1
      }
    }
  }

  return next_cells
}

func countNeighbors(x int, y int) int {
  count := 0
  count += getCell(x-1, y-1) + getCell(x, y-1) + getCell(x+1, y-1)
  count += getCell(x-1, y) + getCell(x+1, y)
  count += getCell(x-1, y+1) + getCell(x, y+1) + getCell(x+1, y+1)
  return count
}

func getCell(x int, y int) int {
  if x < 0 {
    x += yoko
  }
  x = x % yoko

  if y < 0 {
    y += tate
  }
  y = y % tate

  return cells[y][x]
}

func isNextAlive(x int, y int) bool {
  count := countNeighbors(x, y)

  if count == 3 {
    return true
  }

  if cells[y][x] == 1 && (count == 2 || count == 3) {
    return true
  }

  return false
}

func display() {
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()

  for i := 0; i < tate; i++ {
    for j := 0; j < yoko; j++ {
      if cells[i][j] == 1 {
        print("o")
      } else {
        print(" ")
      }
    }
    print("\n")
  }
}
