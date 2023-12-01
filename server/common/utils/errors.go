package utils

type ArgumentError string

func (a ArgumentError) Error() string {
	return "argument error: " + string(a)
}

type ValueError string

func (a ValueError) Error() string {
	return "value invalid: " + string(a)
}

type PathError string

func (p PathError) Error() string {
	return "path error: " + string(p)
}

type FileError string

func (f FileError) Error() string {
	return "file error: " + string(f)
}

type InitializeError string

func (i InitializeError) Error() string {
	return "initialize error: " + string(i)
}

type DataError string

func (m DataError) Error() string {
	return string(m)
}
