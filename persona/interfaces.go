package persona

type Imprimible interface{
	ImprimirInfo()
}

type Identificable interface {
	NombreCompleto() string
	EsUnEstudiante() bool
}