const app = require("./app");
const db = require("./db");
const { HTTP_PORT } = require("./utils/config");

let server;

app.set('db', db);

async function start() {
	const conn = await db.getConnection();
	await conn.ping();
	conn.release();

	server = app.listen(HTTP_PORT, () => {
		console.log(`Server is running on port ${HTTP_PORT}`);
	});
}

async function shutdown(signal) {
  	console.log(`\n${signal} received. Shutting down gracefully...`);

  	try {
		if (server) {
			await new Promise((resolve, reject) => {
				server.close((err) => (err ? reject(err) : resolve()));
			});
			console.log("✅ HTTP server closed");
		}

		if (db) {
			await db.end();
			console.log("✅ Database pool closed");
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