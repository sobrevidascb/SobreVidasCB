package Models

type ACS struct{
	Id int 
	Nome string 
	CPF string
	Nascimento int
	Id_Gen int 
	Telefone int
	Email string
	Mae string
	INE int 
	CBO int 
	CNS int
	CNES int 
	Sts int
}

type Paciente struct{
	Id int 
	Nome string
	CPF string 
	Nascimento int 
	Id_Gen int 
	Telefone int
	Email string
	Mae string
	UF string
	Bairro string
	CEP int
	Municipio string
	Num_casa int 
	Complemento string 
	Tabagista int
	Etilista int
	Absteista int
	Deficiente int 
	M40 int 
	ACS_responsavel int
	Situacao string
	Idade int
}