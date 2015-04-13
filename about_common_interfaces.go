package go_koans

import "bytes"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		in.WriteTo(out)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		in.WriteTo(out)
		out.Truncate(5) // unsure if this is what they wanted me to use or use io.CopyN

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
