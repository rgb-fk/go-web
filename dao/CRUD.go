package dao

func init() {
	gloableRedisHelper.init()
}

const TODO = "TODO"

func GetToDo_EI() (string, error) {
	return gloableRedisHelper.getToDo_EI()
}
func GetToDo_E() (string, error) {
	return gloableRedisHelper.getToDo_E()
}
func GetToDo_I() (string, error) {
	return gloableRedisHelper.getToDo_I()
}
func GetToDo_Other() (string, error) {
	return gloableRedisHelper.getToDo_Other()
}
