// Ringover webhook services
// Handle all business logic here
// these files are being used by ./controllers to pass on data

const { insertIntoDatabase } = require("../repository/webhook.repository");
const logger = require("../utils/logger");

const messageWebhookService = async (app, data) => {
  // Business logic for message webhook
  logger.debug("Message webhook: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Message", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callRingingWebhookService = async (app, data) => {
  // Business logic for call ringing webhook
  logger.debug("CallRingingEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Ringing", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callAnsweredWebhookService = async (app, data) => {
  // Business logic for call answered webhook
  logger.debug("CallAnsweredEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Answered", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callHangupWebhookService = async (app, data) => {
  // Business logic for call hangup webhook
  logger.debug("CallHangupEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Hangup", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callMissedWebhookService = async (app, data) => {
  // Business logic for call missed webhook
  logger.debug("CallMissedEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Missed", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callVoicemailWebhookService = async (app, data) => {
  // Business logic for call voicemail webhook
  logger.debug("CallVoicemailEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Voicemail", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callCommentsUpdatedWebhookService = async (app, data) => {
  // Business logic for call comments updated webhook
  logger.debug("CallCommentsUpdatedEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Comments Updated", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callTagsUpdatedWebhookService = async (app, data) => {
  // Business logic for call tags updated webhook
  logger.debug("CallTagsUpdatedEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Tags Updated", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callRecordAvailableWebhookService = async (app, data) => {
  // Business logic for call record available webhook
  logger.debug("CallRecordAvailableEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Record Available", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const callVoicemailAvailableWebhookService = async (app, data) => {
  // Business logic for call voicemail available webhook
  logger.debug("CallVoicemailAvailableEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Call Voicemail Available", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const handleContact = async (app, data) => {
  // Business logic for contact webhook
  logger.debug("Contact webhook:" + JSON.stringify(data, null, 2));

  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Contact Call", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }

  const contactData = {
    uuid: "b55f949b-c49b-4354-b10a-c8c4cdbd8690",
    firstname: "Jean-Pierre",
    lastname: "De La Court",
    company: "Ringover",
    url: "https://mycrm.com/client/18192233",
    data: {
      key1: "value1",
      key2: "value2",
      keyN: "valueN"
    },
    is_shared: true
  };

  return contactData;
};

const handleContactSearch = async (app, data) => {
  // Business logic for search webhook
  logger.debug("Search webhook:" + JSON.stringify(data, null, 2));

  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Contact Search", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }

  const searchData = [
    {
      firstname: "Jean-Pierre",
      lastname: "De La Court",
      company: "Ringover",
      url: "https://mycrm.com/client/18192233",
      numbers: [
        {
          number: 33184800000,
          type: "mobile"
        }
      ]
    }
  ];

  return searchData;
};

const handleIvr = async (app, data) => {
  // Business logic for IVR response webhook
  logger.debug("IVRResponseEvent: " + JSON.stringify(data, null, 2));
  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook IVR Response", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }
};

const handleSmartRouting = async (app, data) => {
  // Business logic for smart routing webhook
  logger.debug("SmartRouting webhook:" + JSON.stringify(data, null, 2));

  // Do some database operation
  if (app.get("db")) {
    await insertIntoDatabase("Webhook Smart Routing", data.event, data.resource, data.timestamp, data.data, data.attempt);
  }

  const routingData = {
    name: "redirections",
    dispatch: "ringall",
    max_attempts: 1,
    start_delay: 0,
    is_stay_not_connected: true,
    is_stay_in_call: false,
    is_stay_planned_snoozed: true,
    is_stay_snoozed: false,
    ring_overlap: 0,
    agents: [
      {
        agent_type: "agent_external",
        ring_duration: 25,
        ring_delay: 0,
        order: 1,
        number: 33123456789,
        is_pre_answer: false,
        is_caller_id: true,
        is_head_line: false
      }
    ]
  };

  return routingData;
};

module.exports = {
  messageWebhookService,
  callRingingWebhookService,
  callAnsweredWebhookService,
  callHangupWebhookService,
  callMissedWebhookService,
  callVoicemailWebhookService,
  callCommentsUpdatedWebhookService,
  callTagsUpdatedWebhookService,
  callRecordAvailableWebhookService,
  callVoicemailAvailableWebhookService,
  handleContact,
  handleContactSearch,
  handleIvr,
  handleSmartRouting,
};