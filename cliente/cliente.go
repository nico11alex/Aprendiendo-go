package cliente

type Cliente struct{
	Nombre string
	Tipo string
}

func New(nombre,tipo string)Cliente

func (c Cliente) ObtenerDescuento()float64

func (c Cliente) String()string