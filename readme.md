Dados imobiliários – Identificação de Oportunidades de Compra de Terrenos no Brasil 
| Análise da vantagem locacional na aquisição.

OBJETIVO

Fornecer dados que sejam úteis na tomada de decisão com relação a escolha de um terreno para negócio, com considerações de preço, área, localização e dados da população da região. Visando um apoio na decisão de encontrar um local adequado que una localização e perfil populacional para uma boa assertividade de negócio.

FONTES

Usando como base os dados da CAIXA (gov,br). Arquivo único de 30 mil linhas com dados de imóveis a venda por todo o Brasil.

Fonte - https://venda-imoveis.caixa.gov.br/sistema/download-lista.asp?fbclid=PAAabNBSteRzWnlOm7lklDYhzUNSwELumnmtQZSG9akMS8rcsJZt4e1tXqewI_aem_AXAsw0s4TBt7UDU4Zublv_QlpXRkhix0xTWKKBKE0U1BP6jH1yUl5hgFrNBS0W2ZGtc

Para análise, será utilizado os dados apenas de tipo Terreno, pensando no contexto empresarial na inserção do negócio. Os campos utilizados foram: UF, CIDADE, BAIRRO, ENDEREÇO, PREÇO, FINANCIÁVEL (Sim/Não) e AREA.

Para dados relacionados ao local, utilizarei como base o CENSO 2022 – IBGE, com dados como rendimento médio per cápita, relação de sexos, percentual da população com ensino superior, total da populaçao e o crescimento médio (taxa).

Fonte - https://censo2022.ibge.gov.br/panorama/mapas.html?tema=rend_dom_per_capita&recorte=N3
(Coleta feita indicador por indicador em tabelas separadas)


TRATAMENTO/CONSIDERAÇÕES

Considerar apenas linhas do tipo Terreno.
Coleta dos dados nas planilhas do IBGE devem ser únicas por UF.


ESTRUTURA DO DATA WAREHOUSE 
    
TABELAS 
  UF_INFO 

 ID -  UF - RENDA_MEDIA - PERCENT_ENSINO_SUP - IDADE_MEDIA – POPULACAO - HOMEM_PARA_100_MULHERES - TX_CRESC_ANUAL_POP NUMERIC 

 IMOVEL_INFO

ID - UF – CIDADE - BAIRRO - ENDERECO – PRECO - FINANCIAVEL - AREA 


FERRAMENTAS

Banco de Dados  - PostgresSql
Interface de BI – Power BI
Linguagem utilizada para realização do processo ETL – Golang

Tabelas para replicação

CREATE TABLE UF_INFO (
  ID INTEGER GENERATED ALWAYS AS IDENTITY,
  UF VARCHAR(2) NOT NULL,
  RENDA_MEDIA NUMERIC,
  PERCENT_ENSINO_SUP NUMERIC,
  IDADE_MEDIA NUMERIC,
  POPULACAO NUMERIC,
  HOMEM_PARA_100_MULHERES NUMERIC,
  TX_CRESC_ANUAL_POP NUMERIC 
);


CREATE TABLE IMOVEL_INFO (
  ID INTEGER GENERATED ALWAYS AS IDENTITY,
  UF VARCHAR(2) NOT NULL,
  CIDADE VARCHAR,
  BAIRRO VARCHAR,
  ENDERECO VARCHAR,
  PRECO NUMERIC,
  FINANCIAVEL VARCHAR(3),
  AREA NUMERIC 
);

