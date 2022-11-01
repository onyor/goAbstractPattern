package main

type Nike struct {
}

func (a *Nike) makeProduct() IProduct {
	return &NikeProduct{
		Product: Product{
			logo: "nike",
			size: 10,
		},
	}
}
