package main

func main() {
	path := GetPath()
	file := GetFile(path)
	questions := GetQuestions(file)
	AskIfUserIsReady()
	AskQuestions(questions)
}
