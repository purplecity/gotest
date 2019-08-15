package main

import "gotest/sql/mysql"


func hehe() {
	go func() {
		mysql.AddOneRecord(&mysql.AdminUsers{Uid:"666"})
	}()
	return

}
func main() {
	go hehe()


}
