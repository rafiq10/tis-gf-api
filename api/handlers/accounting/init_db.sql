USE [TIS_FinanzasGloabal]


IF OBJECT_ID(N'dbo.accounting_FinancialStatements', N'U') IS NOT NULL  
   DROP TABLE [dbo].[accounting_FinancialStatements]; 
   
IF OBJECT_ID(N'dbo.accounting_cmb_FinStatementTypes', N'U') IS NOT NULL  
   DROP TABLE [dbo].[accounting_cmb_FinStatementTypes];  


--------------------------------------
-- accounting_cmb_FinStatementTypes --
--------------------------------------

CREATE TABLE [dbo].[accounting_cmb_FinStatementTypes](
	[reportType] [nchar](25) NOT NULL,
	[shortName] [nchar](3) NOT NULL
 CONSTRAINT [PK_accounting_cmb_FinStatementTypes_1] PRIMARY KEY CLUSTERED 
(
	[shortName] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, FILLFACTOR = 90) ON [PRIMARY]
) ON [PRIMARY]

  insert into [dbo].[accounting_cmb_FinStatementTypes] (reportType,shortName) 
  values
  ('Balance','BLC'),('Cash Flow Statement','CFL'),('Profit and Loss Statement','P_L')

--------------------------------------
--- accounting_FinancialStatements ---
--------------------------------------

CREATE TABLE [dbo].[accounting_FinancialStatements](
	[country] [nchar](3) NOT NULL,
	[reportType] [nchar](3) NOT NULL,
	[reportYear] [int] NOT NULL,
	[reportMonth] [int] NOT NULL,
	[accNum] [varchar](50) NOT NULL,
	[amount] [decimal](18, 2) NOT NULL,
 CONSTRAINT [PK_accounting_FinancialStatements] PRIMARY KEY CLUSTERED 
(
	[country] ASC,
	[reportType] ASC,
	[reportYear] ASC,
	[reportMonth] ASC,
	[accNum] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, FILLFACTOR = 90) ON [PRIMARY]
) ON [PRIMARY]


ALTER TABLE [dbo].[accounting_FinancialStatements]  WITH CHECK ADD  CONSTRAINT [fk_country_accounting_FinancialStatements] FOREIGN KEY([country])
REFERENCES [dbo].[tblCountriesEsp] ([countryESP])

ALTER TABLE [dbo].[accounting_FinancialStatements] CHECK CONSTRAINT [fk_country_accounting_FinancialStatements]

ALTER TABLE [dbo].[accounting_FinancialStatements]  WITH CHECK ADD  CONSTRAINT [fk_reportType_accounting_FinancialStatements] FOREIGN KEY([reportType])
REFERENCES [dbo].[accounting_cmb_FinStatementTypes] ([shortName])

ALTER TABLE [dbo].[accounting_FinancialStatements] CHECK CONSTRAINT [fk_reportType_accounting_FinancialStatements]

ALTER TABLE [dbo].[accounting_FinancialStatements]  WITH CHECK ADD  CONSTRAINT [fk_reportYear_accounting_FinancialStatements] FOREIGN KEY([reportYear])
REFERENCES [dbo].[tblYears] ([myYear])

ALTER TABLE [dbo].[accounting_FinancialStatements] CHECK CONSTRAINT [fk_reportYear_accounting_FinancialStatements]

ALTER TABLE [dbo].[accounting_FinancialStatements]  WITH CHECK ADD  CONSTRAINT [CHK_reportMonth_accounting_FinancialStatements] CHECK  (([reportMonth]>=(0) AND [reportMonth]<=(12)))

ALTER TABLE [dbo].[accounting_FinancialStatements] CHECK CONSTRAINT [CHK_reportMonth_accounting_FinancialStatements]