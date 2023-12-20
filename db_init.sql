CREATE TABLE Equipment (
                             id SERIAL NOT NULL,
                             engine TEXT NOT NULL,
                             color TEXT NOT NULL,
                             transmission TEXT NOT NULL,
                             body TEXT NOT NULL,
                             PRIMARY KEY (id)
);

CREATE TABLE Supplier (
                            id SERIAL NOT NULL,
                            brand TEXT NOT NULL,
                            country TEXT NOT NULL,
                            city TEXT NOT NULL,
                            address TEXT NOT NULL,
                            PRIMARY KEY (id)
);

CREATE TABLE Client (
                          id SERIAL NOT NULL,
                          login TEXT NOT NULL,
                          password TEXT NOT NULL,
                          surname TEXT NOT NULL,
                          name TEXT NOT NULL,
                          status TEXT NOT NULL,
                          PRIMARY KEY (id)
);

CREATE TABLE Merchant (
                            id SERIAL PRIMARY KEY ,
                            name TEXT,
                            country TEXT NOT NULL,
                            city TEXT NOT NULL,
                            address TEXT NOT NULL
);

CREATE TABLE Car (
                       id SERIAL NOT NULL,
                       name TEXT NOT NULL,
                       is_new BOOLEAN NOT NULL,
                       brand_id INTEGER NOT NULL,
                       equipment_id INTEGER NOT NULL,
                       PRIMARY KEY (id),
                       FOREIGN KEY(brand_id) REFERENCES Supplier (id) ON DELETE CASCADE,
                       FOREIGN KEY(equipment_id) REFERENCES Equipment (id) ON DELETE CASCADE
);

CREATE TABLE Catalog (
                           id SERIAL NOT NULL,
                           merchant_id INTEGER NOT NULL,
                           price INTEGER NOT NULL,
                           sale INTEGER NOT NULL,
                           product_id INTEGER NOT NULL,
                           PRIMARY KEY (id),
                           FOREIGN KEY(merchant_id) REFERENCES Merchant (id) ON DELETE CASCADE,
                           FOREIGN KEY(product_id) REFERENCES Car (id) ON DELETE CASCADE
);

CREATE TABLE "Order" (
                         id SERIAL NOT NULL,
                         catalog_pos INTEGER NOT NULL,
                         created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                         delivered_at TIMESTAMP WITHOUT TIME ZONE,
                         client_id INTEGER NOT NULL,
                         PRIMARY KEY (id),
                         FOREIGN KEY(catalog_pos) REFERENCES Catalog (id) ON DELETE CASCADE,
                         FOREIGN KEY(client_id) REFERENCES Client (id) ON DELETE CASCADE
);

--suppliers
INSERT INTO Supplier ("brand", "country", "city", "address") VALUES ('Mercedes', 'Germany', 'Berlin', 'stuttgart');
INSERT INTO Supplier ("brand", "country", "city", "address") VALUES ('Audi', 'Germany', 'Ingolstadt', 'audi street');
INSERT INTO Supplier ("brand", "country", "city", "address") VALUES ('BMW', 'Germany', 'Munich', 'bmw avenue');
INSERT INTO Supplier ("brand", "country", "city", "address") VALUES ('Volkswagen', 'Germany', 'Wolfsburg', 'vw street');

--equipment
INSERT INTO Equipment ("engine", "color", "transmission", "body") VALUES ('diesel', 'black', 'automatic', 'coupe');
INSERT INTO Equipment ("engine", "color", "transmission", "body") VALUES ('gasoline', 'white', 'manual', 'sedan');
INSERT INTO Equipment ("engine", "color", "transmission", "body") VALUES ('gasoline', 'black', 'manual', 'sedan');
INSERT INTO Equipment ("engine", "color", "transmission", "body") VALUES ('electric', 'yellow', 'automatic', 'coupe');
INSERT INTO Equipment ("engine", "color", "transmission", "body") VALUES ('diesel', 'orange', 'manual', 'hatchback');

--cars
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('GLE', True, 1, 3);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('A4', True, 2, 2);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('X5', False, 3, 3);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('Corolla', True, 4, 4);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('C-Class', False, 1, 1);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('A6', True, 2, 5);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('3 Series', True, 3, 2);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('Golf', False, 4, 1);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('GLC', True, 1, 4);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('Q7', False, 2, 3);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('X4', True, 3, 5);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('Passat', True, 4, 2);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('CLS', False, 1, 4);
INSERT INTO Car ("name", "is_new", "brand_id", "equipment_id") VALUES ('A5', True, 2, 1);

--merchant
INSERT INTO Merchant ("name", "country", "city", "address") VALUES ('MajorAuto', 'Germany', 'Berlin', 'stuttgart');
INSERT INTO Merchant ("name", "country", "city", "address") VALUES ('Rolf', 'Germany', 'Berlin', 'stuttgart');
INSERT INTO Merchant ("name", "country", "city", "address") VALUES ('Genser', 'Germany', 'Berlin', 'stuttgart');



--catalog
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (1, 100, 10, 1);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (2, 150, 20, 4);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (3, 200, 15, 2);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (1, 120, 5, 10);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (2, 180, 0, 5);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (3, 250, 3, 6);
INSERT INTO Catalog ("merchant_id", "price", "sale", "product_id") VALUES (1, 130, 2, 7);


--clients
INSERT INTO Client ("login", "password", "surname", "name", "status") VALUES ('peter', '1', 'Petrov', 'Petr', 'Regular');
INSERT INTO Client ("login", "password", "surname", "name", "status") VALUES ('sm', '2', 'Smith', 'John', 'Regular');
INSERT INTO Client ("login", "password", "surname", "name", "status") VALUES ('ko', '3', 'Kozlov', 'Alexander', 'VIP');

--orders
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (3, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 1);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (3, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 1);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (7, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 2);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (7, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 3);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (1, '2022-05-06 12:33:44', NULL,  2);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (6, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 1);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (3, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 2);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (2, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 1);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (5, '2022-05-06 12:33:44', '2022-06-06 12:33:44', 2);
INSERT INTO "Order" ("catalog_pos", "created_at", "delivered_at", "client_id") VALUES (5, '2022-05-06 12:33:44', NULL, 1);
