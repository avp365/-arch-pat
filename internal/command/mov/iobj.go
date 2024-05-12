package mov

type ObjInterface interface {
	GetParameter(string) interface{}
	SetParameter(string, interface{}) error
}
