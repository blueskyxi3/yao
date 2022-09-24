package table

import (
	"strings"

	"github.com/yaoapp/gou"
)

// Export process

func exportProcess() {
	gou.RegisterProcessHandler("yao.table.setting", processSetting)
	gou.RegisterProcessHandler("yao.table.xgen", processXgen)
	gou.RegisterProcessHandler("yao.table.component", processComponent)
	gou.RegisterProcessHandler("yao.table.search", processSearch)
	gou.RegisterProcessHandler("yao.table.get", processGet)
	gou.RegisterProcessHandler("yao.table.find", processFind)
	gou.RegisterProcessHandler("yao.table.save", processSave)
	gou.RegisterProcessHandler("yao.table.create", processCreate)
	gou.RegisterProcessHandler("yao.table.insert", processInsert)
	gou.RegisterProcessHandler("yao.table.update", processUpdate)
	gou.RegisterProcessHandler("yao.table.updatewhere", processUpdateWhere)
	gou.RegisterProcessHandler("yao.table.updatein", processUpdateIn)
	gou.RegisterProcessHandler("yao.table.delete", processDelete)
	gou.RegisterProcessHandler("yao.table.deletewhere", processDeleteWhere)
	gou.RegisterProcessHandler("yao.table.deletein", processDeleteIn)
}

func processXgen(process *gou.Process) interface{} {
	return nil
}

func processComponent(process *gou.Process) interface{} {
	return nil
}

func processSetting(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Setting.MustExec(process)
}

func processSearch(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Search.MustExec(process)
}

func processGet(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Get.MustExec(process)
}

func processSave(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Save.MustExec(process)
}

func processCreate(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Create.MustExec(process)
}

func processFind(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Find.MustExec(process)
}

func processInsert(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Insert.MustExec(process)
}

func processUpdate(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Update.MustExec(process)
}

func processUpdateWhere(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.UpdateWhere.MustExec(process)
}

func processUpdateIn(process *gou.Process) interface{} {
	process.ValidateArgNums(3)
	tab := MustGet(process)
	ids := strings.Split(process.ArgsString(1), ",")
	process.Args[1] = gou.QueryParam{
		Wheres: []gou.QueryWhere{
			{Column: tab.Layout.Primary, OP: "in", Value: ids},
		},
	}
	return tab.Action.UpdateIn.MustExec(process)
}

func processDelete(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.Delete.MustExec(process)
}

func processDeleteWhere(process *gou.Process) interface{} {
	tab := MustGet(process)
	return tab.Action.DeleteWhere.MustExec(process)
}

func processDeleteIn(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	tab := MustGet(process)
	ids := strings.Split(process.ArgsString(1), ",")
	process.Args[1] = gou.QueryParam{
		Wheres: []gou.QueryWhere{
			{Column: tab.Layout.Primary, OP: "in", Value: ids},
		},
	}
	return tab.Action.DeleteIn.MustExec(process)
}