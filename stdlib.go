package tinylisp

func setupStdlib(env *Environment) {
	// IO
	_ = env.Define(Token{Identifier, "println", -1, nil}, &Println{})

	// Math
	_ = env.Define(Token{Identifier, "+", -1, nil}, &Addition{})
	_ = env.Define(Token{Identifier, "-", -1, nil}, &Subtraction{})
	_ = env.Define(Token{Identifier, "*", -1, nil}, &Multiplication{})
	_ = env.Define(Token{Identifier, "/", -1, nil}, &Division{})

	// Collection
	_ = env.Define(Token{Identifier, "first", -1, nil}, &First{})
	_ = env.Define(Token{Identifier, "rest", -1, nil}, &Rest{})
	_ = env.Define(Token{Identifier, "add", -1, nil}, &Add{})
	_ = env.Define(Token{Identifier, "len", -1, nil}, &Len{})
}