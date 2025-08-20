const app = require("./app");
const { HTTP_PORT } = require("./utils/config");

let server;

async function start() {
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