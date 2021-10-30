package main

import{
	"net/http"
}

func main()
{
	println("Hello World")
	http.ListenAndServe(":8080", nil)
}
}

func say(message string)
{
	println("Change")
}
