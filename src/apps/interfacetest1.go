package main

import "lib/publib/github.com/wonderivan/logger"

type DbHander interface {
	Insert()
	Update(name,id string)
	Select() interface{}
	Select1()
	PrintMember()
}

type DbUser struct {
	userName string
	userId   string
}

func (this *DbUser) Insert() {
	logger.Debug("insert DbUser")
}
func (this *DbUser) Update(name ,id string) {
	this.userId = id
	this.userName = name
	logger.Debug("Update DbUser")
}
func (this *DbUser) Select() interface{} {
	dbusr := this
	logger.Debug("Select DbUser")
	return dbusr
}
func (this *DbUser) Select1()  {
	this.userName = "yyyyyyyyyy"
	this.userId = "99999999"
	logger.Debug("Select1 DbUser\\`")
}

func (this *DbUser) PrintMember() {
	logger.Debug("PrintM DbUser userName=%v userId=%v",this.userName,this.userId)
}

/*
type DbRoleInfo struct {
	roleName string
	roleId   string
}

func (this DbRoleInfo) Insert () {
	logger.Debug("Insert DbRole")
}
func (this DbRoleInfo) Update (name,id string) {
	this.roleName = name
	this.roleId = id
	logger.Debug("Update DbRole")
}
func (this DbRoleInfo) Select() interface{} {
	dbrole := this
	logger.Debug("Select DbRole")
	return dbrole
}
func (this DbRoleInfo) PrintMember() {
	logger.Debug("PrintM DbRole roleName=%v roleId=%v",this.roleName,this.roleId)
}*/

func DbInsert (dbhander DbHander) {
	dbhander.Insert()
}
func DbUpdate (dbhander DbHander,name,id string) {
	dbhander.Update(name,id)
}
func DbSelect (dbhander DbHander) interface{}{
	return dbhander.Select()
}
func DbSelect1 (dbhander DbHander) {
	dbhander.Select1()
}

func main() {
	dbuser := &DbUser{"陈贵华","AB053259"}
	//dbrole := DbRoleInfo{"普通员工","1"}
	var dbselectusr DbUser
	//var dbselectrol DbRoleInfo

	dbuser.PrintMember()
	DbInsert(dbuser)
	dbuser.PrintMember()
	dbselectusr = DbSelect(dbuser).(DbUser)	/*强制类型转换*/
	DbSelect1(dbuser)
	dbselectusr.PrintMember()
	DbUpdate(dbuser,"h00000000","11111111111")
	dbselectusr = DbSelect(dbuser).(DbUser)

	dbselectusr.PrintMember()

	/*dbrole.PrintMember()
	dbselectrol = DbSelect(dbrole).(*DbRoleInfo)
	dbselectrol.PrintMember()*/
}
