create database ecom;
use ecom;

create table PERMISIONS(
    id int not null primary key auto_increment,
    permission varchar(255) unique not null
);

create table ROLES(
    id int not null primary key auto_increment,
    role varchar(255) unique not null
);

create table PERMISIONS_ROLES(
    permission_id int not null,
    role_id int not null,
    primary key(permission_id,role_id),
    foreign key(permission_id) references PERMISIONS(id),
    foreign key(role_id) references ROLES(id)
);

create table USERS(
    id int not null auto_increment,
    email varchar(255) not null,
    firstName varchar(255) not null,
    lastName varchar(255) not null,
    password varchar(255) not null,
    role_id int not null,
    primary key(id),
    foreign key(role_id) references ROLES(id)
);

insert into ROLES(id, role) values(1, "Admin");
insert into ROLES(id, role) values(2, "Seller");
insert into ROLES(id, role) values(3, "Customer");

insert into PERMISIONS(id, permission) values (1, "Create");
insert into PERMISIONS(id, permission) values (2, "Read");
insert into PERMISIONS(id, permission) values (3, "Update");
insert into PERMISIONS(id, permission) values (4, "Delete");

-- Admin
insert into PERMISIONS_ROLES(permission_id, role_id) values (1,1); -- Create
insert into PERMISIONS_ROLES(permission_id, role_id) values (2,1); -- Read
insert into PERMISIONS_ROLES(permission_id, role_id) values (3,1); -- Update
insert into PERMISIONS_ROLES(permission_id, role_id) values (4,1); -- Delete

-- Seller
insert into PERMISIONS_ROLES(permission_id, role_id) values (2,2); -- Read
insert into PERMISIONS_ROLES(permission_id, role_id) values (3,2); -- Update

-- Customer
insert into PERMISIONS_ROLES(permission_id, role_id) values (2,3); -- Read