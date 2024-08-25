create database ecom;
use ecom;

create table PERMISSIONS(
    id int not null primary key auto_increment,
    permission varchar(255) unique not null
);

create table ROLES(
    id int not null primary key auto_increment,
    role varchar(255) unique not null
);

create table PERMISSIONS_ROLES(
    permissionId int not null,
    roleId int not null,
    primary key(permissionId,roleId),
    foreign key(permissionId) references PERMISSIONS(id),
    foreign key(roleId) references ROLES(id)
);

create table USERS(
    id int not null auto_increment,
    email varchar(255) not null,
    firstName varchar(255) not null,
    lastName varchar(255) not null,
    password varchar(255) not null,
    roleId int not null,
    createdAt timestamp not null,
    primary key(id),
    foreign key(roleId) references ROLES(id)
);

create table ADDRESSES(
    id int not null auto_increment,
    userId int not null,
    addressType enum('billing','shipping') not null,
    address text not null,
    city varchar(100) not null,
    state varchar(100) not null,
    country varchar(100) not null,
    zipCode varchar(20) not null,
    primary key(id),
    foreign key(userId) references USERS(id)
  );

create table CATEGORIES(
    id int not null auto_increment,
    name varchar(100) not null,
    description text,
    paraentId int,
    primary key(id),
    foreign key(paraentId) references CATEGORIES(id)
);

create table PRODUCTS(
    id int not null auto_increment,
    name varchar(255) not null,
    description text not null,
    price decimal(10,2) not null,
    stock int not null,
    categoryId int not null,
    image varchar(255),
    createdAt timestamp not null,
    updatedAt timestamp not null,
    check (price >= 0 AND stock >= 0),
    primary key(id),
    foreign key(categoryId) references CATEGORIES(id)
);

create table ORDERS(
    id int not null auto_increment,
    userId int not null,
    orderDate date not null,
    status enum('pending','shipped','completed','canceled') not null,
    totalPrice decimal(10,2) not null,
    check (totalPrice >= 0),
    primary key(id),
    foreign key(userId) references USERS(id)
);

create table ORDER_ITEMS(
    id int not null auto_increment,
    orderId int not null,
    productId int not null,
    quantity int not null,
    price decimal(10,2) not null,
    check (quantity >= 0 AND price >= 0),
    primary key(id),
    foreign key(orderId) references ORDERS(id),
    foreign key(productId) references PRODUCTS(id)
);

create table PAYMENTS(
    id int not null auto_increment,
    orderId int not null,
    paymentMethod varchar(100) not null,
    amount decimal(10,2) not null,
    paymentDate date not null,
    status enum('pending','completed','failed') not null,
    check (amount >= 0),
    primary key(id),
    foreign key(orderId) references ORDERS(id)
);
insert into ROLES(id, role) values(1, "Admin");
insert into ROLES(id, role) values(2, "Seller");
insert into ROLES(id, role) values(3, "Customer");

insert into PERMISSIONS(id, permission) values (1, "Create");
insert into PERMISSIONS(id, permission) values (2, "Read");
insert into PERMISSIONS(id, permission) values (3, "Update");
insert into PERMISSIONS(id, permission) values (4, "Delete");

-- Admin
insert into PERMISSIONS_ROLES(permissionId, roleId) values (1,1); -- Create
insert into PERMISSIONS_ROLES(permissionId, roleId) values (2,1); -- Read
insert into PERMISSIONS_ROLES(permissionId, roleId) values (3,1); -- Update
insert into PERMISSIONS_ROLES(permissionId, roleId) values (4,1); -- Delete

-- Seller
insert into PERMISSIONS_ROLES(permissionId, roleId) values (2,2); -- Read
insert into PERMISSIONS_ROLES(permissionId, roleId) values (3,2); -- Update

-- Customer
insert into PERMISSIONS_ROLES(permissionId, roleId) values (2,3); -- Read