DROP TABLE IF EXISTS statistic;
CREATE TABLE IF NOT EXISTS statistic
(
    id      serial  NOT NULL PRIMARY KEY,
    date    date NOT NULL,
    views   int NOT NULL,
    clicks  int NOT NULL,
    cost    numeric NOT NULL,
    cpc     numeric NOT NULL,
    cpm     numeric NOT NULL
);
