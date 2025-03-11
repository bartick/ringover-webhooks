Ringover Webhook API Integration with FastAPI
==============================================

Description
-----------
This project provides a complete implementation of a webhook integration using FastAPI to handle Ringover events.
It includes endpoints for various types of webhook events such as call events, contact events, contact search events,
smart routing events, webhook messages, and IVR events.

Each endpoint:
  - Verifies the HMAC signature to ensure authenticity.
  - Validates the incoming JSON payload using Pydantic models.
  - Logs the event details.
  - Stores events in a MySQL database using SQLAlchemy.
  - Returns sample responses based on the Ringover documentation.

Features
--------
- FastAPI endpoints for multiple Ringover webhook events.
- HMAC signature verification to secure incoming requests.
- Payload validation with Pydantic.
- Event storage in a MySQL database.
- Sample responses for smart routing, contact, and contact search events.

Prerequisites
-------------
- Python 3.8+
- MySQL Database (optional)
- pip (https://pip.pypa.io/en/stable/)

Installation
------------
1. Clone the repository:
```
git clone <repository-url>
```

3. Navigate to the project directory:
```
cd <project-directory>/python
```

4. Install the dependencies:
```
pip install -r requirements.txt
```

Configuration MYSQL (if needeed)
-------------------
Set the following environment variables or create a .env file:

  - RINGOVER_SECRET: Your secret key used for verifying webhook signatures.
  - DATABASE_URL: Your MySQL connection string, for example:
      mysql+pymysql://user:password@localhost/dbname

Running the Application
-----------------------
Start the application using Uvicorn:

```   
uvicorn main:app --host 0.0.0.0 --port 8000 --reload
```
or 
```   
python app.py
```

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
This project is licensed under the MIT License. See the LICENSE file for details.

Contributing
------------
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

Author
------
Made with Love by Ringover for his customers
