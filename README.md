# SobreVidasCB

Este repositório contém a base para criação de um sistema web baseado em JavaScript, HTML, CSS e GO, com acesso à banco de dados PostgreSQL.

**link para o pitch:https://drive.google.com/file/d/1uEEF0VcxihOtw-IRBb_zalInT-RUrgPH/view?usp=drivesdk**

**link para o vídeo de explicação dos detalhes da implementação:https://drive.google.com/file/d/1Rx8aNKM6c3W-Q8Ie4HXxM0Ju345dl7h8/view?usp=drivesdk**

Com a implementação desse produto, espera-se **auxiliar o trabalho do ACS** na **identificação** do grupo de risco para o câncer de boca, de modo que se consiga **aumentar o número de detecções precoces** da doença e evitar, dessa forma, os malefícios do diagnóstico tardio, como as repercussões negativas do pós-operatório e o aumento dos custos no tratamento. Essa solução almeja permitir o **combate eficaz** ao câncer de boca e ser **financeiramente eficiente**, já que poupa recursos que seriam utilizados no tratamento em estágios avançados da doença. Por fim, a sua implementação contribui para uma transição demográfica saudável e de boa qualidade de vida, especialmente para o público masculino maior de quarenta anos.


[![My Skills](https://skillicons.dev/icons?i=golang,js,html,css,postgres)](https://skillicons.dev)

## Acesso ao banco de dados
   1. Iniciar o pgAdmin
   2. Clicar em > Server... e digitar sua senha
   3. Clicar em Database > Registe um novo banco de dados no nome "SobreVidasCB"
  
## Como rodar
   1. Entre no seu VSCode e clone este repositório.
   2. Abra a pasta "Functions", prossiga para a pasta "DB" e abra o arquivo "Database.go"
   3. Na função "OpenConnection()" apenas altere a senha (coloque a mesma senha usada no pgAdmin)
   4. Abra o terminal e execute o comando `go run .\main.go`
   5. Se tudo der certo será exibida uma mensagem: Server rodando na porta 8080
   6. Está no ar! http://localhost:8080

## Legenda das pastas 
   1. Models: CRUD do servidor
   2. Handlers: CRUD do usuário
   3. Template: arquivos HTML, CSS e JS
   4. Config: configurações do banco de dados
----
