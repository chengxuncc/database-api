/*
Navicat PGSQL Data Transfer

Source Server         : dwx
Source Server Version : 100400
Source Host           : 172.19.153.61:5432
Source Database       : postgres
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 100400
File Encoding         : 65001

Date: 2019-03-08 18:13:09
*/


-- ----------------------------
-- Table structure for conmmodity
-- ----------------------------
DROP TABLE IF EXISTS "public"."conmmodity";
CREATE TABLE "public"."conmmodity" (
"Cno" char(8) COLLATE "default",
"cname" varchar(60) COLLATE "default",
"Cquantity" int8,
"Csupplier" varchar(60) COLLATE "default",
"Cpp" float8,
"Csp" float8,
"Cdis" float4
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for employees
-- ----------------------------
DROP TABLE IF EXISTS "public"."employees";
CREATE TABLE "public"."employees" (
"Eno" char(4) COLLATE "default" NOT NULL,
"Ename" varchar(40) COLLATE "default" NOT NULL,
"Eage" int2 NOT NULL,
"Egender" char(2) COLLATE "default" NOT NULL,
"Eposition" varchar(8) COLLATE "default" NOT NULL,
"Esalary" float8 NOT NULL,
" Econtact" varchar(40) COLLATE "default" NOT NULL,
"Esite" varchar(100) COLLATE "default" NOT NULL,
"Eentrytime" date NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for OUT
-- ----------------------------
DROP TABLE IF EXISTS "public"."OUT";
CREATE TABLE "public"."OUT" (
"Wno" char(4) COLLATE "default",
"cno" char(8) COLLATE "default",
"Cname" varchar(60) COLLATE "default",
"Outq" int8,
"Cquantity" int8,
"Otime" date
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for PUT
-- ----------------------------
DROP TABLE IF EXISTS "public"."PUT";
CREATE TABLE "public"."PUT" (
"Wno" char(4) COLLATE "default",
"cno" char(8) COLLATE "default",
"cname" varchar(60) COLLATE "default",
"Putq" int8,
"Cquantity" int8,
"Ptime" date
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for shiyan
-- ----------------------------
DROP TABLE IF EXISTS "public"."shiyan";
CREATE TABLE "public"."shiyan" (
"jinjia" int8 NOT NULL,
"shoujia" int8 NOT NULL,
"geshu" int8
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for Supplier
-- ----------------------------
DROP TABLE IF EXISTS "public"."Supplier";
CREATE TABLE "public"."Supplier" (
"Sno" char(4) COLLATE "default" NOT NULL,
"Sname" varchar(60) COLLATE "default" NOT NULL,
"Cno" char(8) COLLATE "default" NOT NULL,
"Cname" varchar(60) COLLATE "default" NOT NULL,
"Scontact" varchar(60) COLLATE "default" NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Table structure for warehouse
-- ----------------------------
DROP TABLE IF EXISTS "public"."warehouse";
CREATE TABLE "public"."warehouse" (
"Wno" char(4) COLLATE "default" NOT NULL,
"Cno" char(8) COLLATE "default" NOT NULL,
"Cname" varchar(60) COLLATE "default" NOT NULL,
"Cquantity" int8 NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Triggers structure for table PUT
-- ----------------------------
CREATE TRIGGER "nihai" AFTER INSERT ON "public"."PUT"
FOR EACH ROW
EXECUTE PROCEDURE "auditlogfunc"();

-- ----------------------------
-- Primary Key structure for table Supplier
-- ----------------------------
ALTER TABLE "public"."Supplier" ADD PRIMARY KEY ("Sno");
