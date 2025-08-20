// Entry point of the application.

const app = require("./app");
const { HTTP_PORT } = require("./utils/config");
const logger = require("./utils/logger");

let server;

async function start() {
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

		process.exit(0);
	} catch (err) {
		logger.error("❌ Error during shutdown:", err);
		process.exit(1);
	}
}

["SIGINT", "SIGTERM"].forEach((sig) => {
  	process.on(sig, () => shutdown(sig));
});

start();