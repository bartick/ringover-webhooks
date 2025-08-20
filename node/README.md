# Ringover Webhook API Integration (Express)

This project is an **Express.js app** that handles **Ringover webhook events**.
It verifies, validates, logs, and stores events in **MySQL**.

---

## ✨ Features

* ✅ Multiple Express endpoints for Ringover events (calls, contacts, IVR, smart routing, messages).
* 🔐 Request verification with **HS512 signature**.
* 🧩 Payload validation using **Zod**.
* 💾 Event storage with **MySQL** (`mysql2`).
* 📋 Sample responses included (based on Ringover docs).

---

## 🔧 Prerequisites

* Node.js (v20+)
* (Optional) MySQL server
* A Ringover account with API access

---

## 🚀 Installation

```bash
# Clone repo
git clone https://github.com/ringover/ringover-webhooks.git

# Move into project
cd node

# Install dependencies
npm install
```

---

## ⚙️ Configuration

Create a `.env` file with the following:

### HTTP Port

```env
HTTP_PORT=3000
```

### Webhook Secrets

```env
RINGOVER_SECRET_WEBHOOKS_CALL=your_secret
RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL=your_secret
RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH=your_secret
RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE=your_secret
```

### MySQL (optional)

```env
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=ringover_webhooks
DB_PORT=3306
```

### Suppress Logs (optional)

```env
NODE_ENV=production
```

---

## 🗄️ Database Setup (Optional)

1. Start MySQL.
2. Create a database and user:

   ```sql
   CREATE DATABASE ringover_webhooks;
   CREATE USER 'ringover_user'@'%' IDENTIFIED BY 'ringoverpassword';
   GRANT ALL PRIVILEGES ON ringover_webhooks.* TO 'ringover_user'@'%';
   FLUSH PRIVILEGES;
   ```
3. Run the schema:

   ```bash
   mysql -u ringover_user -p ringover_webhooks < src/db/mydb.sql
   ```

---

## ▶️ Running

```bash
# Start app
npm start

# Start with MySQL
npm run start:mysql

# Start with nodemon
npm run watch

# Nodemon + MySQL
npm run watch:mysql

# Suppress logs (production mode)
NODE_ENV=production npm start
```

---

## 🧪 Testing

```bash
# Run Jest tests
npm test
```

Use **Postman** or `curl` to send requests to endpoints.

---

## 📡 Endpoints

### Call Events

* `/ringover/call/voicemail_available`
* `/ringover/call/record_available`
* `/ringover/call/tags_updated`
* `/ringover/call/comments_updated`
* `/ringover/call/voicemail`
* `/ringover/call/missed`
* `/ringover/call/hangup`
* `/ringover/call/answered`
* `/ringover/call/ringing`
* `/ringover/webhook-message`

### Contact Events

* `/ringover/contact` – Returns contact details

### Contact Search
* `/ringover/contact-search-event` – Returns matching contacts

### Smart Routing and IVR

* `/ringover/smart_routing` – Returns routing rules & agents
* `/ringover/webhook-ivr`

---

## 📜 License

BSD License. See [LICENSE](https://github.com/ringover/ringover-webhooks/blob/main/LICENSE).

---

## 🤝 Contributing

PRs and issues welcome! 🚀

---

## ❤️ Author

Made with love by **Ringover** for our customers.
