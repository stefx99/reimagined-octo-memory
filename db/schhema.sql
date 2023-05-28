/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     5/20/2023 2:32:12 PM                         */
/*==============================================================*/


drop table if exists COWORKINGPROSTOR;

drop table if exists KORISNIK;


drop table if exists REZERVACIJE;

/*==============================================================*/
/* Table: COWORKINGPROSTOR                                      */
/*==============================================================*/
create table COWORKINGPROSTOR
(
   ID_MESTA             int not null  comment '',
   CENA                 float not null  comment '',
   primary key (ID_MESTA)
);

/*==============================================================*/
/* Table: KORISNIK                                              */
/*==============================================================*/
create table KORISNIK
(
   ID_KORISNIKA         varchar(20) not null  comment '',
   IME                  varchar(20) not null  comment '',
   PREZIME              varchar(35) not null  comment '',
   TELEFON              varchar(11) not null  comment '',
   EMAIL                varchar(40) not null  comment '',
   primary key (ID_KORISNIKA)
);

/*==============================================================*/
/* Table: REZERVACIJE                                           */
/*==============================================================*/
create table REZERVACIJE
(
   ID_KORISNIKA         varchar(20) not null  comment '',
   ID_MESTA             int not null  comment '',
   ID_REZERVACIJE       int not null  comment '',
   DATUM                date not null  comment '',
   primary key (ID_KORISNIKA, ID_MESTA, ID_REZERVACIJE)
);

alter table REZERVACIJE add constraint FK_REZERVAC_KORISNIK__KORISNIK foreign key (ID_KORISNIKA)
      references KORISNIK (ID_KORISNIKA) on delete restrict on update restrict;

alter table REZERVACIJE add constraint FK_REZERVAC_REZERVACI_COWORKIN foreign key (ID_MESTA)
      references COWORKINGPROSTOR (ID_MESTA) on delete restrict on update restrict;
