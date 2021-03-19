DROP TABLE IF EXISTS statistic;
CREATE TABLE IF NOT EXISTS statistic
(
    id     serial NOT NULL PRIMARY KEY,
    date   date   NOT NULL,
    views  int     DEFAULT 0,
    clicks int     DEFAULT 0,
    cost   numeric DEFAULT 0,
    cpc    numeric DEFAULT 0,
    cpm    numeric DEFAULT 0
);
