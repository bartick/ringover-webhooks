const request = require("supertest");
const app = require("../app");
const db = require("./../db");
const { RINGOVER_SECRET_WEBHOOKS_CALL, RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL, RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH, RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE } = require('../utils/config');

const jwt = require('jsonwebtoken');

describe("Call Event", () => {

    const ringoverAuthorization = jwt.sign({ foo: "bar" }, RINGOVER_SECRET_WEBHOOKS_CALL, { algorithm: "HS512" });

    it("Message sent Event", async () => {
        const res = await request(app)
            .post("/ringover/webhook-message")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "sent",
                "resource": "sms",
                "timestamp": 0,
                "data": {
                    "id": "1234567-234567",
                    "message_id": 1234567,
                    "conversation_id": 234567,
                    "time": 0,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "body": "My SMS content...",
                    "is_internal": false,
                    "is_collaborative": false,
                    "user_id": 1234567890987654400
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Call Ringing Event", async () => {
        const res = await request(app)
            .post("/ringover/call/ringing")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "ringing",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "id": "57622f20-a020-4f9e-814c-62a6123412fa",
                    "call_id": 1234567890987654400,
                    "channel_id": 123456789,
                    "start_time": 1554823493.762305,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "user_id": 12854321,
                    "is_internal": true,
                    "is_anonymous": false,
                    "is_ivr": true,
                    "ivr_data": {
                        "number": "33101020304",
                        "scenario_name": "Opened",
                        "ivr_name": "myIVR"
                    },
                    "user": {
                        "user_id": 123456789,
                        "firstname": "Jean",
                        "lastname": "Dupont",
                        "email": "jean.dupont@ringover.com",
                        "photo": "https://cdn.ringover.com/img/users/default.jpg"
                    },
                    "status": "ringing"
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    });

    it("Call Answered Event", async () => {
        const res = await request(app)
            .post("/ringover/call/answered")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "answered",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "id": "57622f20-a020-4f9e-814c-62a6123412fa",
                    "call_id": 1234567890987654400,
                    "channel_id": 123456789,
                    "status": "answered",
                    "start_time": 1554823400.762305,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "user_id": 12854321,
                    "is_internal": true,
                    "is_anonymous": false,
                    "is_ivr": true,
                    "ivr_data": {
                        "number": "33101020304",
                        "scenario_name": "Opened",
                        "ivr_name": "myIVR"
                    },
                    "user": {
                        "user_id": 123456789,
                        "firstname": "Jean",
                        "lastname": "Dupont",
                        "email": "jean.dupont@ringover.com",
                        "photo": "https://cdn.ringover.com/img/users/default.jpg"
                    }
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Call Hangup Event", async () => {
        const res = await request(app)
            .post("/ringover/call/hangup")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "hangup",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "id": "57622f20-a020-4f9e-814c-62a6123412fa",
                    "call_id": 1234567890987654400,
                    "channel_id": 123456789,
                    "start_time": 1554823400.762305,
                    "hangup_time": 1554823442.762305,
                    "duration_in_seconds": 32,
                    "record": "https://cdr.ringover.com/record/myrecord.mp3",
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "user_id": 12854321,
                    "is_internal": true,
                    "is_anonymous": false,
                    "is_ivr": true,
                    "ivr_data": {
                        "number": "33101020304",
                        "scenario_name": "Opened",
                        "ivr_name": "myIVR"
                    },
                    "user": {
                        "user_id": 123456789,
                        "firstname": "Jean",
                        "lastname": "Dupont",
                        "email": "jean.dupont@ringover.com",
                        "photo": "https://cdn.ringover.com/img/users/default.jpg"
                    }
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Call Missed Event", async () => {
        const res = await request(app)
            .post("/ringover/call/missed")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "missed",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "id": "57622f20-a020-4f9e-814c-62a6123412fa",
                    "call_id": 1234567890987654400,
                    "start_time": 1554823400.762305,
                    "hangup_time": 1554823442.762305,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "user_id": 12854321,
                    "is_internal": true,
                    "is_anonymous": false,
                    "is_ivr": true,
                    "ivr_data": {
                        "number": "33101020304",
                        "scenario_name": "Opened",
                        "ivr_name": "myIVR"
                    },
                    "ivr": {
                        "name": "StandardFacile",
                        "number": "33677887788",
                        "scenario": {
                            "scenario_id": 123456789,
                            "ivr_id": 123456789,
                            "name": "MonSuperScenario",
                            "scenario_type": "voicemail"
                        }
                    },
                    "status": "missed",
                    "reason": "OUT_PLANNING"
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Call Voicemail Event", async () => {
        const res = await request(app)
            .post("/ringover/call/voicemail")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "voicemail",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "id": "57622f20-a020-4f9e-814c-62a6123412fa",
                    "call_id": 1234567890987654400,
                    "start_time": 1554823400.762305,
                    "answered_time": 1554823412.762305,
                    "hangup_time": 1554823442.762305,
                    "duration_in_seconds": 32,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "user_id": 12854321,
                    "is_internal": true,
                    "is_anonymous": false,
                    "is_ivr": true,
                    "ivr_data": {
                    "number": "33101020304",
                    "scenario_name": "Opened",
                    "ivr_name": "myIVR"
                    }
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Comment Updated Event", async () => {
        const res = await request(app)
            .post("/ringover/call/comments_updated")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "comments_updated",
                "resource": "aftercall",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "channel_id": 123456789,
                    "tags": [
                        "tag1",
                        "tag2"
                    ],
                    "comments": "string"
                },
                "attempt": 2
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Tag Updated Event", async () => {
        const res = await request(app)
            .post("/ringover/call/tags_updated")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "tags_updated",
                "resource": "aftercall",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "tags": [
                        "tag1"
                    ]
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Record Available Event", async () => {
        const res = await request(app)
            .post("/ringover/call/record_available")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "record_available",
                "resource": "aftercall",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "record_link": "string",
                    "record_duration": "string"
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Voicemail Available Event", async () => {
        const res = await request(app)
            .post("/ringover/call/voicemail_available")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "voicemail_available",
                "resource": "aftercall",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "voicemail_link": "string",
                    "voicemail_duration": "string"
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })
})

describe("Contact Call", () => {
    const ringoverAuthorization = jwt.sign({ foo: "bar" }, RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL, { algorithm: "HS512" });

    it("Contact Call Event", async () => {
        const res = await request(app)
            .post("/ringover/contact")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "contact",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304"
                }
        });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({
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
        });
    })

    it("Contact Call Invalid body", async () => {
        const res = await request(app)
            .post("/ringover/contact")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "contact",
                "ressource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304"
                }
            });


        console.log(res.body);
        expect(res.statusCode).toBe(400);
        expect(res.body).toEqual({
            error: [
                {
                    expected: 'string',
                    code: 'invalid_type',
                    path: ['resource'],
                    message: 'Invalid input: expected string, received undefined'
                }
            ]
        });
    })
})

describe("Contact Search", () => {
    const ringoverAuthorization = jwt.sign({ foo: "bar" }, RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH, { algorithm: "HS512" });

    it("Contact Search Event", async () => {
        const res = await request(app)
            .post("/ringover/contact-search-event")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "contact",
                "resource": "search",
                "timestamp": 1554823493.762305,
                "data": {
                    "query_search": "Jean",
                    "user_id": 123456
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual([
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
        ]);
    })

    it("Contact Authorization Fail", async () => {
        const invalidSignature = jwt.sign({ foo: "bar" }, RINGOVER_SECRET_WEBHOOKS_CALL, { algorithm: "HS512" });
        const res = await request(app)
            .post("/ringover/contact")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", invalidSignature) // mock
            .send({
                "event": "contact",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304"
                }
            });

        expect(res.statusCode).toBe(401);
        expect(res.body).toEqual({ error: "Unauthorized" });
    })

    it("Contact X-Ringover-Webhook-Signature required", async () => {
        const res = await request(app)
            .post("/ringover/contact")
            .set("Content-Type", "application/json")
            .send({
                "event": "contact",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304"
                }
            });

        expect(res.statusCode).toBe(401);
        expect(res.body).toEqual({ error: "Missing X-Ringover-Webhook-Signature header" });
    })
})

describe("Smart Routing and IVR Code", () => {
    const ringoverAuthorization = jwt.sign({ foo: "bar" }, RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE, { algorithm: "HS512" });

    it("IVR Response Code Event", async () => {
        console.log(ringoverAuthorization)
        const res = await request(app)
            .post("/ringover/webhook-ivr")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "ivr_response_code",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "code": 1234,
                    "from_number": "33601020304",
                    "to_number": "33101020304",
                    "direction": "inbound",
                    "call_id": 1234567890987654400
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({});
    })

    it("Smart Routing Event", async () => {
        const res = await request(app)
            .post("/ringover/smart_routing")
            .set("Content-Type", "application/json")
            .set("X-Ringover-Webhook-Signature", ringoverAuthorization) // mock
            .send({
                "event": "routing",
                "resource": "call",
                "timestamp": 1554823493.762305,
                "data": {
                    "call_id": 1234567890987654400,
                    "direction": "inbound",
                    "from_number": "33601020304",
                    "to_number": "33101020304"
                }
            });

        expect(res.statusCode).toBe(200);
        expect(res.body).toEqual({
            "name": "redirections",
            "dispatch": "ringall",
            "max_attempts": 1,
            "start_delay": 0,
            "is_stay_not_connected": true,
            "is_stay_in_call": false,
            "is_stay_planned_snoozed": true,
            "is_stay_snoozed": false,
            "ring_overlap": 0,
            "agents": [
                {
                "agent_type": "agent_external",
                "ring_duration": 25,
                "ring_delay": 0,
                "order": 1,
                "number": 33123456789,
                "is_pre_answer": false,
                "is_caller_id": true,
                "is_head_line": false
                }
            ]
        });
    })
})