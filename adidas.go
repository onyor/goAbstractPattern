package main

type Adidas struct {
}

func (a *Adidas) makeProduct() IProduct {
	return &AdidasProduct{
		Product: Product{
			logo: "adidas",
			size: 10,
		},
	}
}
