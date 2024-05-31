package main
type Status int8

const (
	Error Status = iota -1
	Waiting
	Warming
	Complete
)

//implementing the stringer interface
//func (s Status) String() string{
//	switch s {
//	case Error:
//		return "Error"
//	case Waiting:
//		return "Waiting"
//	case Warming:
//		return "Warming"
//	case Complete:
//		return "Complete"
//	default:
//		return ""
//
//	}
//}

//better way of implementing, takes iota as index
func (s Status) String() string{
	return [...]string{"Error", "Waiting", "Warming", "Complete"}[s+1]
}


func main(){
	println(Error)
	println(Error.String())
}
