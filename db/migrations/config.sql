-- +migrate Up
CREATE TABLE IF NOT EXISTS configurations(
        id                                      VARCHAR PRIMARY KEY CHECK (id ~ '^[0-9]'),
        t                                       VARCHAR NOT NULL,
        d                                       VARCHAR NOT NULL,
        host                                    VARCHAR NOT NULL,
        port                                    VARCHAR NOT NULL,
        db                                      VARCHAR NOT NULL,
        u                                       VARCHAR NOT NULL,
        pass                                    VARCHAR NOT NULL,
        sch                                     VARCHAR NOT NULL,
        vhost                                   VARCHAR,
        created_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- +migrate Up
INSERT INTO configurations(id, t, d, host, port, db, u, pass, sch, vhost) VALUES (0, 'Develop.mr_robot', 'Database.processing', 'localhost', '5432', 'devdb', 'mr_robot', 'secret', 'public', NULL);
-- +migrate Up
INSERT INTO configurations(id, t, d, host, port, db, u, pass, sch, vhost) VALUES (1, 'Test.vpn', 'Rabbit.log', '10.0.5.42', '5671', 'devdb', 'guest', 'guest', 'public', '/');
-- +migrate Down
DROP TABLE IF EXISTS configurations;