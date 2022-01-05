--- SET UP CMOBO AND CONFIG TABLES ---

IF OBJECT_ID(N'dbo.tblCountriesEsp', N'U') IS NOT NULL  
   DROP TABLE [dbo].[tblCountriesEsp];  

IF OBJECT_ID(N'dbo.tblYears', N'U') IS NOT NULL  
   DROP TABLE [dbo].[tblYears]; 


--------------------------------------
-- tblCountriesEsp --
--------------------------------------
CREATE TABLE [dbo].[tblCountriesEsp](
	[countryESP] [nchar](3) NOT NULL,
	[countryBankCode] [nchar](2) NULL,
	[bankCheckDigits] [nchar](2) NULL,
PRIMARY KEY CLUSTERED 
(
	[countryESP] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

insert into [dbo].[tblCountriesEsp] (countryESP)
values('BRA'),('CHI'),('COL'),('ECU'),('ESP'),('PER')


--------------------------------------
--- tblYears ---
--------------------------------------
CREATE TABLE [dbo].[tblYears](
	[myYear] [int] NOT NULL,
 CONSTRAINT [PK_tblYears] PRIMARY KEY CLUSTERED 
(
	[myYear] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

insert into [dbo].[tblYears] (myYear)
values(2015),(2016),(2017),(2018),(2019),(2020),(2021),(2022),(2023),(2024),(2025),(2026),(2027),(2028),(2029)


