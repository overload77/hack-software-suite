package code

func TranslateArithmetic(commandName string) string {
	return ArithmeticCommandsAndWriters[commandName]()
}

// TODO
func TranslateMemory(pushOrPop string, segment string, index string) string {
	return ""
}