package Handlers

import (
	Database "SobreVidasCB-layout/Functions/DB"
	"SobreVidasCB-layout/Functions/Models"
	"database/sql"
)

func InsertPC(data Models.Paciente,resp int)bool{
	db := Database.OpenConnection()
	var teste int 
	db.Query("INSERT INTO pacientes(nome,cpf,nascimento,id_gen,telefone,email,mae,uf,bairro,tabagista,etilista,m40,absenteista,acsresponsavel,situacao,def,idade) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17);",data.Nome,data.CPF,data.Nascimento,data.Id_Gen,data.Telefone,data.Email,data.Mae,data.UF,data.Bairro,data.Tabagista,data.Etilista,data.M40,data.Absteista,data.ACS_responsavel,data.Situacao,data.Deficiente,data.Idade)
	row := db.QueryRow("SELECT * FROM pacientes WHERE cpf=$1",data.CPF)
	err := row.Scan(&teste)


	db.Close()
	return err!=sql.ErrNoRows
}