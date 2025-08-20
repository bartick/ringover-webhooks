// Ringover webhook routes
// These file tells what are routes and requests allowed under /ringover

const router = require('express').Router();

const ringoverController = require('./../controllers/ringover.controller');
const verifySignature = require('./../middlewares/verifySignature');
const { RINGOVER_SECRET_WEBHOOKS_CALL, RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL, RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH, RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE } = require("./../utils/config");


router.post('/webhook-message', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleMessageWebhooks);
router.post('/call/ringing', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallRingingWebhooks);
router.post('/call/answered', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallAnsweredWebhooks);
router.post('/call/hangup', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallHangupWebhooks);
router.post('/call/missed', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallMissedWebhooks);
router.post('/call/voicemail', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallVoicemailWebhooks);
router.post('/call/comments_updated', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallCommentsUpdatedWebhooks);
router.post('/call/tags_updated', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallTagsUpdatedWebhooks);
router.post('/call/record_available', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallRecordAvailableWebhooks);
router.post('/call/voicemail_available', verifySignature(RINGOVER_SECRET_WEBHOOKS_CALL), ringoverController.handleCallVoicemailAvailableWebhooks);

router.post('/contact', verifySignature(RINGOVER_SECRET_WEBHOOKS_CONTACT_CALL), ringoverController.handleContactWebhooks);
router.post('/contact-search-event', verifySignature(RINGOVER_SECRET_WEBHOOKS_CONTACT_SEARCH), ringoverController.handleContactSearchWebhooks);

router.post('/webhook-ivr', verifySignature(RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE), ringoverController.handleIvrWebhooks);
router.post('/smart_routing', verifySignature(RINGOVER_SECRET_SMART_ROUTING_AND_IVR_CODE), ringoverController.handleSmartRoutingWebhooks);

module.exports = router;