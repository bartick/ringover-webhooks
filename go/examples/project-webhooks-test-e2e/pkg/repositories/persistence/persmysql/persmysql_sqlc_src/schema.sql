CREATE TABLE webhook_events (
  id          BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  resource    VARCHAR(255) NOT NULL,
  event       VARCHAR(255) NOT NULL,
  timestamp   TIMESTAMP NOT NULL,
  data        JSON NOT NULL
);