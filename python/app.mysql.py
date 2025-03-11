import os
import hmac
import hashlib
import logging
from datetime import datetime
from fastapi import FastAPI, Request, HTTPException, Header
from pydantic import BaseModel
from sqlalchemy import create_engine, Column, Integer, String, DateTime, JSON
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# Configure logging
logging.basicConfig(level=logging.INFO)
app = FastAPI()

# Secrets and database configuration
RINGOVER_SECRET = os.getenv("RINGOVER_SECRET", "your_secret_here")
DATABASE_URL = os.getenv("DATABASE_URL", "mysql+pymysql://user:password@localhost/dbname")
HTTP_PORT = os.getenv("RINGOVER_HTTP_PORT", "8089")

# Set up SQLAlchemy engine and session
engine = create_engine(DATABASE_URL, echo=True,future=True)
SessionLocal = sessionmaker(bind=engine,future=True)
Base = declarative_base()

# Database model to store all webhook events
class WebhookEventModel(Base):
    __tablename__ = "webhook_events"
    id = Column(Integer, primary_key=True, index=True)
    event = Column(String(255))
    timestamp = Column(String(255))  # ISO formatted timestamp stored as string
    data = Column(JSON)
    created_at = Column(DateTime, default=datetime.utcnow)

Base.metadata.create_all(bind=engine)

# ----------------------
# Pydantic models for the various webhook payloads
# ----------------------

# Generic model for call events (used for multiple endpoints)
class RingoverWebhook(BaseModel):
    event: str
    data: dict
    timestamp: str  # ISO formatted date/time

# Contact event (generic)
class RingoverContactWebhook(BaseModel):
    query: str
    contacts: list
    timestamp: str  # ISO formatted date/time

# Contact search event (specific)
class RingoverContactSearchEvent(BaseModel):
    event: str  # Expected to be "contactSearchEvent"
    query: str
    contacts: list
    timestamp: str  # ISO formatted date/time

# Webhook Message event
class RingoverWebhookMessage(BaseModel):
    event: str  # Expected to be "webhookMessage"
    message: str
    timestamp: str  # ISO formatted date/time

# Webhook IVR event
class RingoverWebhookIVR(BaseModel):
    event: str  # Expected to be "webhookIVR"
    ivr: dict
    timestamp: str  # ISO formatted date/time

# Smart Routing event
class RingoverSmartRoutingEvent(BaseModel):
    event: str  # Expected to be "smart_routing" or similar
    data: dict
    timestamp: str  # ISO formatted date/time

# ----------------------
# Helper function to verify the HMAC signature
# ----------------------
def verify_signature(body: bytes, signature: str) -> bool:
    computed_signature = hmac.new(
        RINGOVER_SECRET.encode(), body, hashlib.sha256
    ).hexdigest()
    return hmac.compare_digest(computed_signature, signature)

# ----------------------
# Helper to store events in DB
# ----------------------
def store_event(event_name: str, timestamp: str, payload: dict):
    with SessionLocal() as db:
        try:
            event_record = WebhookEventModel(
                event=event_name,
                timestamp=timestamp,
                data=payload
            )
            db.add(event_record)
            db.commit()
            db.refresh(event_record)
            logging.info(f"Stored event '{event_name}' with ID: {event_record.id}")
        except Exception as e:
            db.rollback()
            logging.error(f"Error storing event '{event_name}': {e}")

# ----------------------
# Endpoints for call events
# ----------------------
@app.post("/ringover/call/voicemail_available")
async def call_voicemail_available(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_voicemail_available event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/record_available")
async def call_record_available(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_record_available event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/tags_updated")
async def call_tags_updated(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_tags_updated event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/comments_updated")
async def call_comments_updated(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_comments_updated event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/voicemail")
async def call_voicemail(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_voicemail event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/missed")
async def call_missed(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_missed event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/hangup")
async def call_hangup(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_hangup event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/answered")
async def call_answered(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_answered event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/call/ringing")
async def call_ringing(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        event_data = RingoverWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received call_ringing event at {event_data.timestamp}")
    store_event(event_data.event, event_data.timestamp, payload)
    return {"status": "ok"}

# ----------------------
# Endpoints for contact events
# ----------------------
@app.post("/ringover/contact")
async def contact_event(request: Request, x_ringover_signature: str = Header(None)):
    """
    Sample Response (from Ringover documentation):
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
    """
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        contact_event = RingoverContactWebhook.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received contact event with query '{contact_event.query}' at {contact_event.timestamp}")
    store_event("contact", contact_event.timestamp, payload)
    
    # Return sample response for the contact webhook
    return {
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

@app.post("/ringover/contact-search-event")
async def contact_search_event(request: Request, x_ringover_signature: str = Header(None)):
    """
    Sample Response (from Ringover documentation):
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
    """
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        search_event = RingoverContactSearchEvent.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received contact search event with query '{search_event.query}' at {search_event.timestamp}")
    store_event(search_event.event, search_event.timestamp, payload)
    
    # Return sample response for the contact search event
    return [
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

# ----------------------
# Endpoints for webhook message and IVR events
# ----------------------
@app.post("/ringover/webhook-message")
async def webhook_message(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        message_event = RingoverWebhookMessage.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received webhookMessage event at {message_event.timestamp}: {message_event.message}")
    store_event(message_event.event, message_event.timestamp, payload)
    return {"status": "ok"}

@app.post("/ringover/webhook-ivr")
async def webhook_ivr(request: Request, x_ringover_signature: str = Header(None)):
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        ivr_event = RingoverWebhookIVR.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received webhookIVR event at {ivr_event.timestamp} with data: {ivr_event.ivr}")
    store_event(ivr_event.event, ivr_event.timestamp, payload)
    return {"status": "ok"}

# ----------------------
# Endpoint for Smart Routing events
# ----------------------
@app.post("/ringover/smart_routing")
async def smart_routing_event(request: Request, x_ringover_signature: str = Header(None)):
    """
    Sample Response (from Ringover documentation):
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
    """
    body = await request.body()
    if x_ringover_signature is None:
        raise HTTPException(status_code=400, detail="Signature header missing")
    if not verify_signature(body, x_ringover_signature):
        raise HTTPException(status_code=400, detail="Invalid signature")
    try:
        payload = await request.json()
        smart_event = RingoverSmartRoutingEvent.parse_obj(payload)
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Invalid payload: {e}")

    logging.info(f"Received smart routing event '{smart_event.event}' at {smart_event.timestamp}")
    logging.info(f"Data: {smart_event.data}")
    store_event(smart_event.event, smart_event.timestamp, payload)
    
    # Return sample response for the smart routing event
    return {
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

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=int(HTTP_PORT), reload=True)
