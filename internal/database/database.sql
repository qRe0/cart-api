CREATE TABLE "carts"("id" SERIAL NOT NULL);
ALTER TABLE
    "carts" ADD PRIMARY KEY("id");
CREATE TABLE "items"(
                        "id" SERIAL NOT NULL,
                        "cart_id" INTEGER NOT NULL,
                        "product" TEXT NOT NULL,
                        "quantity" INTEGER NOT NULL,
                        UNIQUE (cart_id, product)
);
ALTER TABLE
    "items" ADD PRIMARY KEY("id");
ALTER TABLE
    "items" ADD CONSTRAINT "items_cart_id_foreign" FOREIGN KEY("cart_id") REFERENCES "carts"("id");