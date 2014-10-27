package main

const prog = `var x = 3a
var y = 4
var z = x + y

func main() {
	{
		print(x);
	}
	;
}

print(z) // some comment
`
