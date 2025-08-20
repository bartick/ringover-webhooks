// Load environment variables from .env file

require('dotenv').config({
    quiet: true,
});

module.exports = {
    // production?
    NODE_ENV: process.env.NODE_ENV || "development",

    // ringover secrets
    RINGOVER_SECRET_WEBHOOKS_CALL: process.env.RINGOVER_SECRET_WEBHOOKS_CALL || "default_webhooks_secret",
    RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL: process.env.RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL || "default_webhooks_contact_secret",
    RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH: process.env.RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH || "default_webhooks_search_secret",
    RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE: process.env.RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE || "default_smart_routing_and_ivr_code",
    
    // webhook start port
    HTTP_PORT: process.env.HTTP_PORT || 3000,

    // Database
    DB_HOST: process.env.DB_HOST || "localhost",
    DB_USER: process.env.DB_USER || "root",
    DB_PASSWORD: process.env.DB_PASSWORD || "rootpassword",
    DB_NAME: process.env.DB_NAME || "mydb",
    DB_PORT: process.env.DB_PORT || 3306
}