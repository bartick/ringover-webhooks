// Zod schemas for Ringover webhook events.
// These schemas might change in future (data of writing: 20th Aug, 2025)
// For more info refer to: https://developer.ringover.com/#tag/webhook

const { z } = require("zod");

exports.IVRResponseEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        code: z.number(),
        from_number: z.string(),
        to_number: z.string(),
        direction: z.enum(["inbound", "outbound"])
    })
});

exports.MessageEvent = z.object({
    event: z.enum(["received", "sent"]),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        id: z.string(),
        message_id: z.number().or(z.string()),
        conversation_id: z.number().or(z.string()),
        time: z.number(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        body: z.string(),
        is_internal: z.boolean(),
        is_collaborative: z.boolean(),
        user_id: z.number().or(z.string()),
    }),
    attempt: z.number()
});

exports.SmartRoutingEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
    }),
});

exports.ContractEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
    }),
});

exports.ContractSearchEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        query_search: z.string(),
        user_id: z.number().or(z.string()),
    }),
});

exports.CallRingingEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.looseObject({
        id: z.string(),
        call_id: z.number().or(z.string()),
        channel_id: z.number().or(z.string()),
        start_time: z.number(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        user_id: z.number().or(z.string()),
        is_internal: z.boolean(),
        is_anonymous: z.boolean(),
        is_ivr: z.boolean(),
        ivr_data: z.object({
            number: z.string(),
            scenario_name: z.string(),
            ivr_name: z.string(),
        }).optional().nullable(),
        user: z.object({
            user_id: z.number().or(z.string()),
            firstname: z.string(),
            lastname: z.string(),
            email: z.string(),
            photo: z.string()
        }),
        status: z.string(),
    }),
    attempt: z.number()
});

exports.CallAnsweredEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        id: z.string(),
        call_id: z.number().or(z.string()),
        channel_id: z.number().or(z.string()),
        status: z.string(),
        start_time: z.number(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        user_id: z.number().or(z.string()),
        is_internal: z.boolean(),
        is_anonymous: z.boolean(),
        is_ivr: z.boolean(),
        ivr_data: z.object({
            number: z.string(),
            scenario_name: z.string(),
            ivr_name: z.string(),
        }).optional().nullable(),
        user: z.object({
            user_id: z.number().or(z.string()),
            firstname: z.string(),
            lastname: z.string(),
            email: z.string(),
            photo: z.string()
        }),
        status: z.string(),
    }),
    attempt: z.number()
});

exports.CallHangupEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        id: z.string(),
        call_id: z.number().or(z.string()),
        channel_id: z.number().or(z.string()),
        start_time: z.number(),
        hangup_time: z.number(),
        duration_in_seconds: z.number().optional().nullable(),
        record: z.string().optional().nullable(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        user_id: z.number().or(z.string()),
        is_internal: z.boolean(),
        is_anonymous: z.boolean(),
        is_ivr: z.boolean(),
        ivr_data: z.object({
            number: z.string(),
            scenario_name: z.string(),
            ivr_name: z.string(),
        }).optional().nullable(),
        user: z.object({
            user_id: z.number().or(z.string()),
            firstname: z.string(),
            lastname: z.string(),
            email: z.string(),
            photo: z.string()
        }),
    }),
    attempt: z.number()
});

exports.CallMissedEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        id: z.string(),
        call_id: z.number().or(z.string()),
        start_time: z.number(),
        hangup_time: z.number(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        user_id: z.number().or(z.string()),
        is_internal: z.boolean(),
        is_anonymous: z.boolean(),
        is_ivr: z.boolean(),
        ivr_data: z.object({
            number: z.string(),
            scenario_name: z.string(),
            ivr_name: z.string(),
        }).optional().nullable(),
        status: z.string(),
        reason: z.string(),
    }),
    attempt: z.number()
});

exports.CallVoicemailEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        id: z.string(),
        call_id: z.number().or(z.string()),
        start_time: z.number(),
        answered_time: z.number().nullable().optional(),
        hangup_time: z.number(),
        duration_in_seconds: z.number().optional().nullable(),
        direction: z.enum(["inbound", "outbound"]),
        from_number: z.string(),
        to_number: z.string(),
        user_id: z.number().or(z.string()),
        is_internal: z.boolean(),
        is_anonymous: z.boolean(),
        is_ivr: z.boolean(),
        ivr_data: z.object({
            number: z.string(),
            scenario_name: z.string(),
            ivr_name: z.string(),
        }).optional().nullable(),
    }),
});

exports.CommentUpdatedEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        channel_id: z.number().or(z.string()),
        tags: z.array(z.string()),
        comments: z.string(),
    }),
    attempt: z.number()
});

exports.TagUpdatedEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        tags: z.array(z.string()),
    }),
});

exports.RecordAvailableEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        record_link: z.string().optional().nullable(),
        record_duration: z.string().or(z.number()).optional().nullable(),
    }),
});

exports.VoicemailAvailableEvent = z.object({
    event: z.string(),
    resource: z.string(),
    timestamp: z.number(),
    data: z.object({
        call_id: z.number().or(z.string()),
        voicemail_link: z.string().optional().nullable(),
        voicemail_duration: z.string().or(z.number()).optional().nullable(),
    }),
});