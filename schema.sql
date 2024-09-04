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
    name varchar(100) not null unique,
    description text,
    parentId int,
    primary key(id),
    foreign key(parentId) references CATEGORIES(id)
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

insert into USERS(email, firstName, password, roleId, createdAt) values
("admin@mail.com", "admin", "", "6u6mXJogpVIo1CZ0QkjRrFtv80e5deCGYdwFIsiXbn7Rx7exPdeMFqE", 1, CURRENT_TIMESTAMP());

insert into CATEGORIES (name, description, parentId) values
('Electronics', 'Electronic devices and accessories', NULL),
('Clothing', 'Apparel and fashion items', NULL),
('Books', 'Books and literature', NULL),
('HomeGoods', 'Household items and furnishings', NULL),
('Electronics Smartphones', 'Mobile phones', 1),
('Electronics Computers', 'Desktops, laptops, and components', 1),
('Electronics Gaming', 'Gaming consoles and peripherals', 1),
('Clothing Men', 'Clothing for men', 2),
('Clothing Women', 'Clothing for wamen', 2),
('Clothing Kids', 'Clothing for children', 2),
('Books Fiction', 'Fictional works', 3),
('Books NonFiction', 'Non-fictional works', 3),
('HomeGoods Kitchen', 'Kitchenware and appliances', 4),
('HomeGoods Furniture', 'Furniture for home and office', 4);

insert into PRODUCTS (name, description, price, stock, categoryId, image, createdAt, updatedAt) VALUES
('iPhone 14 Pro', 'Latest iPhone model with Pro features', 1199.99, 50, 2, 'iphone14pro.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Gaming Laptop', 'High-performance gaming laptop', 1599.99, 20, 3, 'gaminglaptop.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Jeans for men', 'Classic denim jeans', 49.99, 100, 6, 'mensjeans.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Dress for women', 'Elegant evening dress', 99.99, 30, 7, 'womensdress.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Science Fiction Novel', 'A thrilling sci-fi adventure', 19.99, 50, 10, 'scifinovel.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Cookbook', 'Recipes for delicious home-cooked meals', 24.99, 15, 13, 'cookbook.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP()),
('Sofa', 'Modern leather sofa', 999.99, 10, 14, 'sofa.jpg',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP());