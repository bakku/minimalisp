package tinylisp

func setupStdlib(env *Environment) {
	// IO
	_ = env.Define(Token{Identifier, "println", -1, nil}, &Println{})

	// Math
	_ = env.Define(Token{Identifier, "+", -1, nil}, &Addition{})
	_ = env.Define(Token{Identifier, "-", -1, nil}, &Subtraction{})
	_ = env.Define(Token{Identifier, "*", -1, nil}, &Multiplication{})
	_ = env.Define(Token{Identifier, "/", -1, nil}, &Division{})
}
