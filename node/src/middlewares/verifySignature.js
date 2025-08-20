const jwt = require("jsonwebtoken");

function verifySignature(secret) {
  return (req, res, next) => {
    try {
      const signature = req.headers["x-ringover-webhook-signature"];

      if (!signature) {
        return res
          .status(401)
          .json({ error: "Missing X-Ringover-Webhook-Signature header" });
      }

      // Verify JWT (HS512 only)
      jwt.verify(
        signature,
        secret,
        { algorithms: ["HS512"] },
        (err, decoded) => {
          if (err) {
            return res.status(401).send({ error: "Unauthorized" });
          }

          // Attach decoded token for downstream handlers
          req.jwtPayload = decoded;

          next(); // ✅ verified
        }
      );
    } catch (err) {
      console.error("Signature verification failed:", err);
      return res.status(500).json({ error: "Internal server error" });
    }
  };
}

module.exports = verifySignature;
