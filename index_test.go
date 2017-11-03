package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

//問３
func Testancestor(t *testing.T) {
	fmt.Println("カテゴリ名を入力してください")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	if text == "ファッション" || text == "家具・インテリア" {
		fmt.Println("先祖カテゴリです")
	} else {
		t.Error("failed")
	}
}

func Testmain(t *testing.T)  {
  stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
  if text == "ファッション" {
    fmt.Println("子孫カテゴリは「ファッション」です")
  } else if text == "レディース" {
    fmt.Println("子孫カテゴリは「レディース／靴」です")
  } else if text == "メンズ" {
    fmt.Println("子孫カテゴリは「メンズ」です")
  } else if text == "靴" {
    fmt.Println("子孫カテゴリは「靴」です")
  } else if text == "家具・インテリア" {
    fmt.Println("子孫カテゴリは「家具・インテリア／収納家具」です")
  } else if text == "収納家具" {
    fmt.Println("子孫カテゴリは「収納家具」です")
  } else {
    t.Error("failed")
  }
}
