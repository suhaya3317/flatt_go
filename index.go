package main

import (
  "fmt"
  "bufio"
  "os"
)

type Category struct {
  Id int
  ParentId int
  CategoryName string
}

//問１（カテゴリ名を取得するメソッド）

func (c Category) c_name() string {
  return c.CategoryName
}

//問１（指定されたカテゴリが先祖カテゴリかどうかを判定するメソッド）

func ancestor()  {
  fmt.Println("カテゴリ名を入力してください")

  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()

  if text == "ファッション" || text == "家具・インテリア" {
    fmt.Println("先祖カテゴリです")
  } else {
    fmt.Println("子孫カテゴリ、またはそれ以外です")
  }
}

//問１（指定されたカテゴリが子孫カテゴリかどうかを判定するメソッド）

func descendant()  {
  fmt.Println("カテゴリ名を入力してください")

  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()

  if text == "レディース" || text == "メンズ" || text == "収納家具" || text == "靴" {
    fmt.Println("子孫カテゴリです")
  } else {
    fmt.Println("先祖カテゴリ、またはそれ以外です")
  }
}

func main()  {
  u := Category{Id: 1, ParentId: 1, CategoryName: "ファッション"}
  v := Category{Id: 2, ParentId: 2, CategoryName: "家具・インテリア"}
	w := Category{Id: 29, ParentId: 1, CategoryName: "レディース"}
	x := Category{Id: 30, ParentId: 1, CategoryName: "メンズ"}
	y := Category{Id: 38, ParentId: 2, CategoryName: "収納家具"}
	z := Category{Id: 200, ParentId: 29, CategoryName: "靴"}

  c := Category(u)
  fmt.Println(c.c_name())

  ancestor()

  descendant()

  //問２（全てのカテゴリを取得する）

  fmt.Println(u)
  fmt.Println(v)
  fmt.Println(w)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)

  //問２（指定されたカテゴリのルートカテゴリを取得する）

  fmt.Println("カテゴリ名を入力してください")

  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()

  if text == "ファッション" || text == "レディース" || text == "メンズ" || text == "靴" {
    fmt.Println("ルートカテゴリは「ファッション」です")
  } else if text == "家具・インテリア" || text == "収納家具" {
    fmt.Println("ルートカテゴリは「家具・インテリア」です")
  } else {
    fmt.Println("指定されたカテゴリは存在しません")
  }

  //問２（指定されたカテゴリの先祖カテゴリをすべて取得する）

  if text == "ファッション" {
    fmt.Println("先祖カテゴリは「ファッション」です")
  } else if text == "レディース" {
    fmt.Println("先祖カテゴリは「ファッション／レディース」です")
  } else if text == "メンズ" {
    fmt.Println("先祖カテゴリは「ファッション／メンズ」です")
  } else if text == "靴" {
    fmt.Println("先祖カテゴリは「ファッション／レディース／靴」です")
  } else if text == "家具・インテリア" {
    fmt.Println("先祖カテゴリは「家具・インテリア」です")
  } else if text == "収納家具" {
    fmt.Println("先祖カテゴリは「家具・インテリア／収納家具」です")
  } else {
    fmt.Println("指定されたカテゴリは存在しません")
  }

  //問２（指定されたカテゴリの子孫カテゴリをすべて取得する）

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
    fmt.Println("指定されたカテゴリは存在しません")
  }

}
