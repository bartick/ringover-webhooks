const { z } = require("zod");

const ringoverService = require("../services/ringover.service");

const { MessageEvent, CallRingingEvent, CallAnsweredEvent, CallHangupEvent, CallMissedEvent, CallVoicemailEvent, CommentUpdatedEvent, TagUpdatedEvent, RecordAvailableEvent, VoicemailAvailableEvent, ContractEvent, ContractSearchEvent, IVRResponseEvent, SmartRoutingEvent } = require('../zod/ringover.zod');
const logger = require("../utils/logger");

const handleMessageWebhooks = async (req, res) => {
  try {
    const data = MessageEvent.parse(req.body);

    await ringoverService.messageWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Message Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid MessageEvent");
    }
  }
}

const handleCallRingingWebhooks = async (req, res) => {
  try {
    const data = CallRingingEvent.parse(req.body);

    await ringoverService.callRingingWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Ringing Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CallRingingEvent");
    }
  }
}

const handleCallAnsweredWebhooks = async (req, res) => {
  try {
    const data = CallAnsweredEvent.parse(req.body);

    await ringoverService.callAnsweredWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Answered Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CallAnsweredEvent");
    }
  }
}

const handleCallHangupWebhooks = async (req, res) => {
  try {
    const data = CallHangupEvent.parse(req.body);

    await ringoverService.callHangupWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Hangup Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CallHangupEvent");
    }
  }
}

const handleCallMissedWebhooks = async (req, res) => {
  try {
    const data = CallMissedEvent.parse(req.body);

    await ringoverService.callMissedWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Missed Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CallMissedEvent");
    }
  }
}

const handleCallVoicemailWebhooks = async (req, res) => {
  try {
    const data = CallVoicemailEvent.parse(req.body);

    await ringoverService.callVoicemailWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Voicemail Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CallVoicemailEvent");
    }
  }
}

const handleCallCommentsUpdatedWebhooks = async (req, res) => {
  try {
    const data = CommentUpdatedEvent.parse(req.body);

    await ringoverService.callCommentsUpdatedWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Comment Updated Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid CommentUpdatedEvent");
    }
  }
}

const handleCallTagsUpdatedWebhooks = async (req, res) => {
  try {
    const data = TagUpdatedEvent.parse(req.body);

    await ringoverService.callTagsUpdatedWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Tag Updated Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid TagUpdatedEvent");
    }
  }
}

const handleCallRecordAvailableWebhooks = async (req, res) => {
  try {
    const data = RecordAvailableEvent.parse(req.body);

    await ringoverService.callRecordAvailableWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Record Available Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid RecordAvailableEvent");
    }
  }
}

const handleCallVoicemailAvailableWebhooks = async (req, res) => {
  try {
    const data = VoicemailAvailableEvent.parse(req.body);

    await ringoverService.callVoicemailAvailableWebhookService(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Call Voicemail Available Webhook error: " + JSON.stringify(err, null, 2));
      res.status(400).send("Invalid VoicemailAvailableEvent");
    }
  }
}

const handleContactWebhooks = async (req, res) => {
  try {
    const data = ContractEvent.parse(req.body);

    const response = await ringoverService.handleContact(req.app, data);

    res.status(200).send(response);
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Contact webhook error: " + JSON.stringify(err, null, 2));
      res.status(500).send("Internal Server Error");
    }
  }
};

const handleContactSearchWebhooks = async (req, res) => {
  try {
    const data = ContractSearchEvent.parse(req.body);

    const response = await ringoverService.handleContactSearch(req.app, data);

    res.status(200).send(response);
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("Contact search webhook error: " + JSON.stringify(err, null, 2));
      res.status(500).send("Internal Server Error");
    }
  }
};

const handleIvrWebhooks = async (req, res) => {
  try {
    const data = IVRResponseEvent.parse(req.body);

    await ringoverService.handleIvr(req.app, data);

    res.status(200).send();
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("IVR webhook error: " + JSON.stringify(err, null, 2));
      res.status(500).send("Internal Server Error");
    }
  }
}

const handleSmartRoutingWebhooks = async (req, res) => {
  try {
    const data = SmartRoutingEvent.parse(req.body);

    const response = await ringoverService.handleSmartRouting(req.app, data);

    res.status(200).send(response);
  } catch (err) {
    if (err instanceof z.ZodError) {
      logger.error("ZodError: " + JSON.stringify(JSON.parse(err.message), null, 2));
      res.status(400).json({ error: JSON.parse(err.message) });
    } else {
      logger.error("SmartRouting webhook error: " + JSON.stringify(err, null, 2));
      res.status(500).send("Internal Server Error");
    }
  }
}

module.exports = {
  handleMessageWebhooks,
  handleCallRingingWebhooks,
  handleCallAnsweredWebhooks,
  handleCallHangupWebhooks,
  handleCallMissedWebhooks,
  handleCallVoicemailWebhooks,
  handleCallCommentsUpdatedWebhooks,
  handleCallTagsUpdatedWebhooks,
  handleCallRecordAvailableWebhooks,
  handleCallVoicemailAvailableWebhooks,
  handleContactWebhooks,
  handleContactSearchWebhooks,
  handleIvrWebhooks,
  handleSmartRoutingWebhooks,
}