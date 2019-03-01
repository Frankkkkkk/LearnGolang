package main

var currentId int

var todos Todos


func init(){
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host "})
	//readFile()
}

//func RepoFindTodo(id int)Todo{
//	for _, t :=range todos{
//		if t.Id ==id{
//			return t
//		}
//	}
//	return Todo{}
//}

func RepoCreateTodo(t Todo)Todo{
	currentId +=1
	t.Id = currentId
	todos = append(todos, t)
	return t
}
//func RepoDestoryTodo(id int)error{
//	for i,t :=range todos{
//		if t.Id == id{
//			todos = append(todos[:1], todos[i+1:]...)
//			return nil
//		}
//	}
//	return fmt.Errorf("Could not find Todo with id of %d to delete",id)
//}