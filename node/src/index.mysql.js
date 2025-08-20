const app = require("./app");
const db = require("./db");
const { HTTP_PORT } = require("./utils/config");
const logger = require("./utils/logger");

let server;

app.set('db', db);

async function start() {
	const conn = await db.getConnection();
	await conn.ping();
	conn.release();

	server = app.listen(HTTP_PORT, () => {
		logger.debug(`Server is running on port ${HTTP_PORT}`);
	});
}

async function shutdown(signal) {
  	logger.debug(`\n${signal} received. Shutting down gracefully...`);

  	try {
		if (server) {
			await new Promise((resolve, reject) => {
				server.close((err) => (err ? reject(err) : resolve()));
			});
			logger.debug("✅ HTTP server closed");
		}

		if (db) {
			await db.end();
			logger.debug("✅ Database pool closed");
		}

		process.exit(0);
	} catch (err) {
		console.error("❌ Error during shutdown:", err);
		process.exit(1);
	}
}

["SIGINT", "SIGTERM"].forEach((sig) => {
  	process.on(sig, () => shutdown(sig));
});

start();