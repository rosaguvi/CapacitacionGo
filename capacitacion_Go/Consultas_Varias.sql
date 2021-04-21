User name,Password,Access key ID,Secret access key,Console login link
Administrator,,,,https://804896286407.signin.aws.amazon.com/console
Lukas,Luk4s4587/*-,AKIA3WZ4OU3DZS3W4FRJ,bu3sP/KXDIEgLn7FNp7yysyUguqdQ405LacvC09J,https://804896286407.signin.aws.amazon.com/console
Marcos,M4rc0s7845/*-,AKIA3WZ4OU3DZZYEMYBF,Ok36yDww0Zo8mUERqDAxVW9fA4HNXUtGc+PPW0WC,https://804896286407.signin.aws.amazon.com/console
DataUSer,D4t4Us3r2356,AKIA3WZ4OU3DY3K3IBX7,L5WvSbNK+OBdm/esRPoCFKmWEilxlFK7TkAFY5Sl,https://804896286407.signin.aws.amazon.com/console

elasticloadbalancing:DescribeLoadBalancer

184.168.131.241

*/
-- Database: ProAdmModPer


Repositorios

cd go/src/Proyecto-example/posts-backend/


cd go/src/github.com/proyectoGo/

-- DROP DATABASE "ProAdmModPer";

CREATE DATABASE "ProAdmModPer"
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Spanish_Colombia.1252'
    LC_CTYPE = 'Spanish_Colombia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE "ProAdmModPer"
    IS 'Base de Datos Para el Proyectos de Administración de Perfiles de Go';
	
-- insertar Roles 

INSERT INTO public.rols(
	created_at, updated_at, name, )
	VALUES 
	(NOW(), NOW() ,'SUPER_ADMIN'),
	(NOW(), NOW() ,'ADMIN'),
	(NOW(), NOW() ,'USER');
	
-- insertar Perfiles

INSERT INTO public.perfils(
	created_at, updated_at, name , usuario_gb)
	VALUES (now(),now(), 'Jefe de Sistemas',1)

-- insertar Modulos

INSERT INTO public.modulos(
	 created_at, updated_at, name, usuario_gb)
	VALUES 
	(now(),now(),'Usuarios',1),
	(now(),now(),'Perfiles',1),
	(now(),now(),'Modulos',1)
	
	
-- insertar Usuarios

INSERT INTO public.usuarios(
	 created_at, updated_at, name, email, password, perfil_id, rol_id , usuario_gb)
	VALUES 
	 (now(), now(),'Juan', 'Juan@personalsoft.com', '$2a$08$nsfezcuZV0lsHJpIV7N2ZOHbbr/.zThlyP17Klm5Ypcjmu/b0PsUe', 1, 1,1)