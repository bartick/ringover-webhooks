// Webhook repository
// Here we are defining the database operations for webhook events

const db = require("../db");

const insertIntoDatabase = async (eventType, event, resource, timestamp, data, attempt) => {
    const sql = `
        INSERT INTO webhook_events 
        (event_type, event, resource, timestamp, attempt, data) 
        VALUES (?, ?, ?, ?, ?, ?)
    `;

    const values = [
        eventType,
        event,
        resource,
        timestamp,
        attempt || null,
        JSON.stringify(data),
    ];

    await db.execute(sql, values);
}

module.exports = {
    insertIntoDatabase
};
