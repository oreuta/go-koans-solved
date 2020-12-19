package go_koans

import "bytes"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // YOU HAVE CREATED out BUT DIDN'T FILL IT WITH ANY DATA.... PLEASE FIX IT
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		assert(out.String() == "hello") // YOU HAVE CREATED out BUT DIDN'T FILL IT WITH ANY DATA.... PLEASE FIX IT
	}
}
