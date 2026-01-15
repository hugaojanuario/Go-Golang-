package d1

type Pessoa struct{
	Nome string
	Idade int
}

func (p *Pessoa) MaiorDeIdade() bool{
	if p.Idade >= 18{
		return true
	}else{
		return false
	}
}