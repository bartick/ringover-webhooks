// Database connection pool
// Create a connection pool for MySQL database

const mysql = require("mysql2/promise");
const { DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT } = require("../utils/config");

const pool = mysql.createPool({
  host: DB_HOST,
  user: DB_USER,
  password: DB_PASSWORD,
  database: DB_NAME,
  port: DB_PORT,
  waitForConnections: true,
  connectionLimit: 1,
  queueLimit: 0,
});

module.exports = pool;
