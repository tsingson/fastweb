package utils

// First returns first argument of a function return
// use it like
//     myutils.First(ab())
func First(args ...interface{}) interface{} {
	if len(args) == 0 {
		return nil
	}
	return args[0]
}

// Last returns first argument of a function return
// use it like
//     myutils.Last(ab())
func Last(args ...interface{}) interface{} {
	if len(args) == 0 {
		return nil
	}
	return args[len(args)]
}

// Pick returns picked argument of a function return
// use it like
//     myutils.Pick(1, ab())
func Pick(index int, args ...interface{}) interface{} {
	if len(args) == 0 {
		return nil
	}

	return args[index]
}

// Slice returns arguments of a function return as a slice
// use it like
//     myutils.Slice(ab())[1]
func Slice(args ...interface{}) []interface{} {
	return args
}
