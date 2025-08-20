Ringover Webhook API Integration with Express
==============================================

Description
-----------
This project provides a complete implementation of a webhook integration using `Express` to handle Ringover events.
It includes endpoints for various types of webhook events such as call events, contact events, contact search events,
smart routing events, webhook messages, and IVR events.

Each endpoint:
  - Verifies the `HS512` signature to ensure authenticity.
  - Validates the incoming JSON payload using `Zod` schemas.
  - Logs the event details.
  - Stores events in a MySQL database using `mysql2`.
  - Returns sample responses based on the Ringover documentation.

Features
--------
- Express endpoints for multiple Ringover webhook events.
- HS512 signature verification to secure incoming requests.
- Payload validation with Zod.
- Event storage in a MySQL database.
- Sample responses for smart routing, contact, and contact search events.

Prerequisites
-------------
- Node.js (version 20 or higher)
- MySQL database (optional)
- Ringover account and API access

Installation
------------
1. Clone the repository:
```
git clone https://github.com/ringover/ringover-webhooks.git
```

3. Navigate to the project directory:
```
cd node
```

4. Install the dependencies:
```
npm install
```

Configuration Main
-------------------
Set the following environment variables or create a .env file:

 - RINGOVER_SECRET_WEBHOOKS_CALL: Your secret key given by ringover
 - RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL: Your secret key given by ringover
 - RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH: Your secret key given by ringover
 - RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE: Your secret key given by ringover
 - HTTP_PORT: Your HTTP port


Configuration MYSQL (optional)
-------------------
Set the following environment variables or create a .env file:

 - DB_HOST: Your database host
 - DB_USER: Your database user
 - DB_PASSWORD: Your database password
 - DB_NAME: Your database name
 - DB_PORT: Your database port

Run MySql (optional)
----------
1. Start your MySQL server.
2. Create a database for the project:
```
CREATE DATABASE <your_database_name>;
```
3. Create a user and grant privileges (replace `user` and `password` with your own values):
```
CREATE USER '<your_database_user>'@'<your_database_host>' IDENTIFIED BY '<your_database_password>';
GRANT ALL PRIVILEGES ON <your_database_name>.* TO '<your_database_user>'@'<your_database_host>';
FLUSH PRIVILEGES;
```
4. Use the database and run the create table command in `<project_root>/node/src/db/mydb.sql`:
```
mysql -u <your_username> -p <your_database_name> < <project_root>/node/src/db/mydb.sql
```

Running the Application
-----------------------
1. Start the application:
```
npm start
```
2. Start the application with mysql:
```
npm run start:mysql
```
3. Start the application using nodemon:
```
npm run watch
```
4. Start the application with mysql using nodemon:
```
npm run watch:mysql
```
5. Run server but do not want the events logs
```
NODE_ENV=production npm start
```
5. Test locally
```
npm test
```
6. Test the webhook endpoints using a tool like Postman or curl.

Endpoints
---------
```
Call Events:
  - /ringover/call/voicemail_available
  - /ringover/call/record_available
  - /ringover/call/tags_updated
  - /ringover/call/comments_updated
  - /ringover/call/voicemail
  - /ringover/call/missed
  - /ringover/call/hangup
  - /ringover/call/answered
  - /ringover/call/ringing
```

Contact Events:
  - /ringover/contact
    Sample Response:
```
{
		"uuid": "b55f949b-c49b-4354-b10a-c8c4cdbd8690",
		"firstname": "Jean-Pierre",
		"lastname": "De La Court",
		"company": "Ringover",
		"url": "https://mycrm.com/client/18192233",
		"data": {
			"key1": "value1",
			"key2": "value2",
			"keyN": "valueN"
		},
		"is_shared": true
}
```
  - /ringover/contact-search-event
    Sample Response:
```
[
	{
		"firstname": "Jean-Pierre",
		"lastname": "De La Court",
		"company": "Ringover",
		"url": "https://mycrm.com/client/18192233",
		"numbers": [
			{
				"number": 33184800000,
				"type": "mobile"
			}
		]	
	}
]
```   
Other Webhook Events:
  - /ringover/webhook-message
  - /ringover/webhook-ivr

Smart Routing:
  - /ringover/smart_routing
    Sample Response:
```
{
	"name": "redirections",
	"dispatch": "ringall",
	"max_attempts": 1,
	"start_delay": 0,
	"is_stay_not_connected": 1,
	"is_stay_in_call": 0,
	"is_stay_planned_snoozed": 1,
	"is_stay_snoozed": 0,
	"ring_overlap": 0,
	"agents": [
		{
			"agent_type": "agent_external",
			"ring_duration": 25,
			"ring_delay": 0,
			"order": 1,
			"number": 33123456789,
			"is_pre_answer": 0,
			"is_caller_id": 1,
			"is_head_line": 0
		}
	]
}      
```
License
-------
This project is licensed under the BSD License. See the LICENSE file for details.

Contributing
------------
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

Author
------
Made with Love by Ringover for his customers