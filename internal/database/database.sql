CREATE TABLE "carts"
(
    "id" SERIAL PRIMARY KEY
);

CREATE TABLE "items"
(
    "id"       SERIAL PRIMARY KEY,
    "cart_id"  INTEGER NOT NULL,
    "product"  TEXT    NOT NULL,
    "quantity" INTEGER NOT NULL,
    UNIQUE (cart_id, product),
    CONSTRAINT "items_cart_id_foreign" FOREIGN KEY ("cart_id") REFERENCES "carts" ("id") ON DELETE CASCADE
);
